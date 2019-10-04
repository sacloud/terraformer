package sacloud

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/accessor"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// StateWaiter リソースの状態が変わるまで待機する
type StateWaiter interface {
	// WaitForState リソースが指定の状態になるまで待つ
	WaitForState(context.Context) (interface{}, error)
	// AsyncWaitForState リソースが指定の状態になるまで待つ
	AsyncWaitForState(context.Context) (compCh <-chan interface{}, progressCh <-chan interface{}, errorCh <-chan error)
}

var (
	// DefaultStatePollTimeout StatePollWaiterでのデフォルトタイムアウト
	DefaultStatePollTimeout = 20 * time.Minute
	// DefaultStatePollInterval StatePollWaiterでのデフォルトポーリング間隔
	DefaultStatePollInterval = 5 * time.Second
)

// StateReadFunc StatePollWaiterにより利用される、対象リソースの状態を取得するためのfunc
type StateReadFunc func() (state interface{}, err error)

// StateCheckFunc StateReadFuncで得たリソースの情報を元に待ちを継続するか判定するためのfunc
//
// StatePollWaiterのフィールドとして設定する
type StateCheckFunc func(target interface{}) (exit bool, err error)

// UnexpectedAvailabilityError 予期しないAvailabilityとなった場合のerror
type UnexpectedAvailabilityError struct {
	// Err エラー詳細
	Err error
}

// Error errorインターフェース実装
func (e *UnexpectedAvailabilityError) Error() string {
	return fmt.Sprintf("resource returns unexpected availability value: %s", e.Err.Error())
}

// UnexpectedInstanceStatusError 予期しないInstanceStatusとなった場合のerror
type UnexpectedInstanceStatusError struct {
	// Err エラー詳細
	Err error
}

// Error errorインターフェース実装
func (e *UnexpectedInstanceStatusError) Error() string {
	return fmt.Sprintf("resource returns unexpected instance status value: %s", e.Err.Error())
}

// StatePollWaiter ポーリングによりリソースの状態が変わるまで待機する
type StatePollWaiter struct {
	// NotFoundRetry Readで404が返ってきた場合のリトライ回数
	//
	// アプライアンスなどの一部のリソースでは作成~起動完了までの間に404を返すことがある。
	// これに対応するためこのフィールドにて404発生の許容回数を指定可能にする。
	NotFoundRetry int

	// ReadFunc 対象リソースの状態を取得するためのfunc
	//
	// TargetAvailabilityを指定する場合はAvailabilityHolderを返す必要がある
	// もしAvailabilityHolderを実装しておらず、かつStateCheckFuncも未指定だった場合はタイムアウトまで完了しないため注意
	ReadFunc StateReadFunc

	// TargetAvailability 対象リソースのAvailabilityがこの状態になった場合になるまで待つ
	//
	// この値を指定する場合、ReadFuncにてAvailabilityHolderを返す必要がある。
	// AvailabilityがTargetAvailabilityとPendingAvailabilityで指定されていない状態になった場合はUnexpectedAvailabilityErrorを返す
	//
	// TargetAvailability(Pending)とTargetInstanceState(Pending)の両方が指定された場合は両方を満たすまで待つ
	// StateCheckFuncとの併用は不可。併用した場合はpanicする。
	TargetAvailability []types.EAvailability

	// PendingAvailability 対象リソースのAvailabilityがこの状態になった場合は待ちを継続する。
	//
	// 詳細はTargetAvailabilityのコメントを参照
	PendingAvailability []types.EAvailability

	// TargetInstanceStatus 対象リソースのInstanceStatusがこの状態になった場合になるまで待つ
	//
	// この値を指定する場合、ReadFuncにてInstanceStatusHolderを返す必要がある。
	// InstanceStatusがTargetInstanceStatusとPendinngInstanceStatusで指定されていない状態になった場合はUnexpectedInstanceStatusErrorを返す
	//
	// TargetAvailabilityとTargetInstanceStateの両方が指定された場合は両方を満たすまで待つ
	//
	// StateCheckFuncとの併用は不可。併用した場合はpanicする。
	TargetInstanceStatus []types.EServerInstanceStatus

	// PendingInstanceStatus 対象リソースのInstanceStatusがこの状態になった場合は待ちを継続する。
	//
	// 詳細はTargetInstanceStatusのコメントを参照
	PendingInstanceStatus []types.EServerInstanceStatus

	// StateCheckFunc ReadFuncで得たリソースの情報を元に待ちを継続するかの判定を行うためのfunc
	//
	// TargetAvailabilityとTargetInstanceStateとの併用は不可。併用した場合panicする
	StateCheckFunc StateCheckFunc

	// Timeout タイムアウト
	Timeout time.Duration // タイムアウト
	// PollInterval ポーリング間隔
	PollInterval time.Duration
}

func (w *StatePollWaiter) validateFields() {
	if w.ReadFunc == nil {
		panic(errors.New("StatePollWaiter has invalid setting: ReadFunc is required"))
	}

	if w.StateCheckFunc != nil && (len(w.TargetAvailability) > 0 || len(w.TargetInstanceStatus) > 0) {
		panic(errors.New("StatePollWaiter has invalid setting: StateCheckFunc and TargetAvailability/TargetInstanceStatus can not use together"))
	}

	if w.StateCheckFunc == nil && len(w.TargetAvailability) == 0 && len(w.TargetInstanceStatus) == 0 {
		panic(errors.New("StatePollWaiter has invalid setting: TargetAvailability or TargetInstanceState must have least 1 items when StateCheckFunc is not set"))
	}
}

func (w *StatePollWaiter) defaults() {

	if w.Timeout == time.Duration(0) {
		w.Timeout = DefaultStatePollTimeout
	}
	if w.PollInterval == time.Duration(0) {
		w.PollInterval = DefaultStatePollInterval
	}
}

// WaitForState リソースが指定の状態になるまで待つ
func (w *StatePollWaiter) WaitForState(ctx context.Context) (interface{}, error) {
	c, p, e := w.AsyncWaitForState(ctx)
	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("WaitForState is canceled")
		case lastState := <-c:
			return lastState, nil
		case <-p:
			// noop
		case err := <-e:
			return nil, err
		}
	}
}

// AsyncWaitForState リソースが指定の状態になるまで待つ
func (w *StatePollWaiter) AsyncWaitForState(ctx context.Context) (compCh <-chan interface{}, progressCh <-chan interface{}, errorCh <-chan error) {

	w.validateFields()
	w.defaults()

	compChan := make(chan interface{})
	progChan := make(chan interface{})
	errChan := make(chan error)

	tick := time.Tick(w.PollInterval)
	bomb := time.After(w.Timeout)

	go func() {
		notFoundCounter := w.NotFoundRetry
		for {
			select {
			case <-ctx.Done():
				errChan <- errors.New("AsyncWaitForState is canceled")
				return
			case <-tick:
				state, err := w.ReadFunc()

				if err != nil {
					if IsNotFoundError(err) {
						notFoundCounter--
						if notFoundCounter > 0 {
							continue
						}
					}
					errChan <- fmt.Errorf("AsyncWaitForState is failed: %s", err)
					return
				}

				exit, err := w.handleState(state)
				if exit {
					compChan <- state
					return
				}

				if err != nil {
					errChan <- err
					return
				}

				if state != nil {
					progChan <- state
				}
			case <-bomb:
				errChan <- errors.New("AsyncWaitForState is timed out")
				return
			}
		}
	}()
	return compChan, progChan, errChan
}

func (w *StatePollWaiter) handleState(state interface{}) (bool, error) {
	if w.StateCheckFunc != nil {
		return w.StateCheckFunc(state)
	}

	availabilityHolder, hasAvailability := state.(accessor.Availability)
	instanceStateHolder, hasInstanceState := state.(accessor.InstanceStatus)

	switch {
	case hasAvailability && hasInstanceState:

		res1, err := w.handleAvailability(availabilityHolder)
		if err != nil {
			return false, err
		}
		res2, err := w.handleInstanceState(instanceStateHolder)
		if err != nil {
			return false, err
		}
		return res1 && res2, nil

	case hasAvailability:
		return w.handleAvailability(availabilityHolder)
	case hasInstanceState:
		return w.handleInstanceState(instanceStateHolder)
	default:
		// どちらのインターフェースも実装していない場合、stateが存在するだけでtrueとする
		return true, nil
	}
}

func (w *StatePollWaiter) handleAvailability(state accessor.Availability) (bool, error) {
	if len(w.TargetAvailability) == 0 {
		return true, nil
	}
	v := state.GetAvailability()
	switch {
	case w.isInAvailability(v, w.TargetAvailability):
		return true, nil
	case w.isInAvailability(v, w.PendingAvailability):
		return false, nil
	default:
		return false, fmt.Errorf("got unexpected value of Availability: got %q", v)
	}
}

func (w *StatePollWaiter) handleInstanceState(state accessor.InstanceStatus) (bool, error) {
	if len(w.TargetInstanceStatus) == 0 {
		return true, nil
	}
	v := state.GetInstanceStatus()
	switch {
	case w.isInInstanceStatus(v, w.TargetInstanceStatus):
		return true, nil
	case w.isInInstanceStatus(v, w.PendingInstanceStatus):
		return false, nil
	default:
		return false, fmt.Errorf("got unexpected value of InstanceState: got %q", v)
	}
}

func (w *StatePollWaiter) isInAvailability(v types.EAvailability, conds []types.EAvailability) bool {
	for _, cond := range conds {
		if v == cond {
			return true
		}
	}
	return false
}

func (w *StatePollWaiter) isInInstanceStatus(v types.EServerInstanceStatus, conds []types.EServerInstanceStatus) bool {
	for _, cond := range conds {
		if v == cond {
			return true
		}
	}
	return false
}
