# Terraformer

このリポジトリは[GoogleCloudPlatform/terraformer](https://github.com/GoogleCloudPlatform/terraformer)のフォークです。  
オリジナルのTerraformerにさくらのクラウドプロバイダー([sacloud/terraform-provider-sakuracloud](https://github.com/sacloud/terraform-provider-sakuracloud))に対応する機能を追加したものです。

基本的な利用方法などについてはオリジナルの[README.md](https://github.com/GoogleCloudPlatform/terraformer)を参照してください。

# インストール

Dockerを利用する方法と実行ファイルをローカルにダウンロードする方法があります。

## Dockerを利用する場合

Dockerを利用することでTerraform/Terraformer/さくらのクラウドプロバイダーが一式入った環境を手軽に利用できます。

```bash
$ docker run -it --rm -v $PWD:/work sacloud/terraformer
```

Dockerイメージを自分でビルドする場合は以下のようにします。

1.  Run `git clone <terraformer repo>`
2.  Build docker image by `docker build -t terraformer .`
3.  Run `docker run -it --rm -v $PWD:/work terraformer`

## 実行ファイルをローカルにダウンロードする場合

### Terraformのインストール

`terraform`コマンドをパスの通った場所に配置します。

### Terraformerのインストール

From source:

1.  Run `git clone <terraformer repo>`
2.  Run `GO111MODULE=on go mod vendor`
3.  Run `go build -v` for all providers OR build with one provider `go run build/main.go {google,aws,azure,kubernetes and etc}`
4.  Run ```terraform init``` against an ```init.tf``` file to install the plugins required for your platform. For example, if you need plugins for the google provider, ```init.tf``` should contain:
```
provider "google" {}
```
Or alternatively

4.  Copy your Terraform provider's plugin(s) to folder
    `~/.terraform.d/plugins/{darwin,linux}_amd64/`, as appropriate.

From Releases:

* Linux

```
curl -LO https://github.com/sacloud/terraformer/releases/download/$(curl -s https://api.github.com/repos/sacloud/terraformer/releases/latest | grep tag_name | cut -d '"' -f 4)/terraformer-linux-amd64.zip
unzip terraformer-linux-amd64.zip
chmod +x terraformer
sudo mv terraformer-linux-amd64 /usr/local/bin/terraformer
```

* MacOS

```
curl -LO https://github.com/sacloud/terraformer/releases/download/$(curl -s https://api.github.com/repos/sacloud/terraformer/releases/latest | grep tag_name | cut -d '"' -f 4)/terraformer-darwin-amd64.zip
unzip terraformer-darwin-amd64.zip
chmod +x terraformer
sudo mv terraformer-darwin-amd64 /usr/local/bin/terraformer
```

* リリース

リリースページから適切なファイルをダウンロードしてください。

### プロバイダーのインストール

以下のフォルダ内にプロバイダーの実行ファイルをコピーしておいてください。

* `~/.terraform.d/plugins/{darwin,linux}_amd64/`


Note:  さくらのクラウドプロバイダーはv1.16.4以降のバージョンが推奨です。

## 使い方

コマンドラインオプション、または環境変数でAPIトークン/シークレットを指定する必要があります。

### コマンドラインオプションでAPIキーを指定する場合

```
$ terraformer import sakuracloud --token=APIトークン --secret=APIシークレット --resource=server,disk,icon
```

### 環境変数でAPIキーを指定する場合

```
$ export SAKURACLOUD_ACCESS_TOKEN=APIトークン
$ export SAKURACLOUD_ACCESS_TOKEN_SECRET=APIシークレット
$ terraformer import sakuracloud --resource=server,disk,icon
```

利用例:

```
 # サーバのみ、IDを指定
 ./terraformer import sakuracloud --resources=server --filter=sakuracloud_server=id1:id2:id4
```

```
 # 対応している全リソースを指定
 ./terraformer import sakuracloud --resources=archive,autoBackup,bridge,cdrom,database,disk,dns,gslb,icon,internet,loadBalancer,mobileGateway,nfs,note,packetFilter,privateHost,proxyLB,server,sim,simpleMonitor,sshKey,switch,vpcRouter
```

### サポートしているリソース

以下のリソースをサポートしています。
(この一覧は `terraformer import sakuracloud list`コマンドでも確認できます)

* `archive` : `sakuracloud_archive`
* `autoBackup` : `sakuracloud_auto_backup`
* `bridge` : `sakuracloud_bridge`
* `cdrom` : `sakuracloud_cdrom`
* `database` : `sakuracloud_database`
* `disk` : `sakuracloud_disk`
* `dns` : `sakuracloud_dns`
* `gslb` : `sakuracloud_gslb`
* `icon` : `sakuracloud_icon`
* `internet` : `sakuracloud_internet`
* `loadBalancer` : `sakuracloud_load_balancer`
* `mobileGateway` : `sakuracloud_mobile_gateway`
* `nfs` : `sakuracloud_nfs`
* `note` : `sakuracloud_note`
* `packetFilter` : `sakuracloud_packet_filter`
* `privateHost` : `sakuracloud_private_host`
* `proxyLB` : `sakuracloud_proxylb`
* `server` : `sakuracloud_server`
* `sim` : `sakuracloud_sim`
* `simpleMonitor` : `sakuracloud_simple_monitor`
* `sshKey` : `sakuracloud_ssh_key`
* `switch` : `sakuracloud_switch`
* `vpcRouter` : `sakuracloud_vpc_router`

## 利用上の注意

### Terraformのバージョン

Terraform v0.12+が必要です。

Terraform v0.11を利用したい場合はv0.7.9をご利用ください。

### サポートしない項目

サーバリソース(`sakuracloud_server`)のディスクの修正関連パラメータなどの入力専用項目はtfファイル/tfstateファイルに出力されませんのでご注意ください。  
出力されない項目は以下のようなものがあります。  
これらは必要に応じてtfファイルの書き換えを行ってください。

#### アーカイブ(`sakuracloud_archive`)

- `archive_file`: アップロードするアーカイブファイル
- `hash`: アーカイブファイルのハッシュ値

#### ISOイメージ(`sakuracloud_cdrom`)

- `iso_image_file`: ISOイメージファイル
- `content`: ISOイメージファイルのコンテンツ
- `content_file_name`: ISOイメージファイルのファイルパス
- `hash`: ハッシュ値

#### アイコン(`sakuracloud_icon`)

- `source`: アイコンのファイルパス
- `base64content`: アイコンファイルのコンテンツ

#### サーバ(`sakuracloud_server`)

- `disk_edit_parameter`: ディスクの修正関連パラメータ

### パブリックリソースのtfファイル間の参照出力

パブリックリソース(`Scope=Shared`)はtfファイルにIDのみ出力され、`${data.terraform_remote_state.xxx}`という参照は出力されません。

対象となるパブリックリソース:

- アーカイブ
- ISOイメージ
- アイコン
- スタートアップスクリプト

これらは必要に応じて以下のようにtfファイルの書き換えを行ってください。

```hcl
# 書き換え前: パブリックアーカイブの場合アーカイブのIDがtfファイルに出力される
resource "sakuracloud_disk" "disk-001-example" {
  name                      = "example"
  source_archive_id         = "123456789012" // パブリックアーカイブの場合、${}での参照とならない
}
```

```hcl
# 書き換え後: IDをデータソース経由などでの参照に書き換える
data "sakuracloud_archive" "ubuntu" {
  os_type = "ubuntu"
}

resource "sakuracloud_disk" "disk-001-example" {
  name                      = "example"
  source_archive_id         = "${data.sakuracloud_archive.ubuntu.id}" // データソースを参照するように書き換え
}
```

### tfファイル上で参照出力されずIDが出力される場合

`terraformer import`実行時の対象リソースの設定で参照元/参照先両方のリソースを指定する必要があります。  
例えばサーバリソース(`sakuracloud_server`)の`icon_id`を参照にしたい場合は以下のようにサーバとアイコンの両方をオプションで指定してください。

```
terraformer import sakuracloud -r server,icon
```

## バージョニングについて

オリジナルのTerraformerのリリースタグごとに対応したsacloud/terraformerをリリースします。

sacloud/terraformでのリリースタグは以下のようなルールで付与します。

    sacloud/<オリジナルのタグ>

もし同一のオリジナルタグに対して複数のバージョンのsacloud/terraformerをリリースする場合は以下のようなルールでタグを付与します。  
(連番は0開始)

    sacloud/<オリジナルのタグ>/patch-<リリースごとの連番>
    
例えば、オリジナルのリリースタグが `0.7.9`の場合、対応するリリースタグは `sacloud/0.7.9`、以降は `sacloud/0.7.9/patch-0`となります。

## License

 `sacloud/terraformer` Copyright 2019 Kazumichi Yamamoto.

  This project is published under [Apache 2.0 License](LICENSE).
  
