# sacloud-upload-image

## Overview

さくらのクラウドへISOイメージのアップロードを行うためのコマンドです。

## Description

さくらのクラウドには[ISOイメージ](http://cloud.sakura.ad.jp/specification/server-disk/)という機能があります。
通常さくらのクラウドでサーバを作成する際はあらかじめ初期設定済みのOSイメージ(アーカイブ)からOSを選択するのですが、自分でISOイメージをアップロードして好きなOSをインストールすることも可能です。

ただ、ISOイメージのアップロードは割と煩雑な作業です。詳細は[こちらのページ](https://help.sakura.ad.jp/app/answers/detail/a_id/2520)に記載されているのですが、

    1) コントロールパネル上でISOファイル用領域を作成
    2) ファイル領域へのFTP接続情報が表示されるため控えておく
    3) FTPS+PASV対応のFTPクライアントからアップロード

という手順が必要になります。

`sacloud-upload-image`を使えばコマンド一発でアップロードが行えるようになります。

## Install

こちらの[リリースページ](https://github.com/yamamoto-febc/sacloud-upload-image/releases/latest)から最新バージョンのバイナリをダウンロードして展開、実行権を付与してください。

以下のプラットフォーム用のバイナリを用意しています。
  * darwin(i386/amd64)
  * linux(i386/amd64)
  * windows(i386/amd64)

## Usage

オプション、イメージのパス、イメージ名を指定するとアップロード実施します。
コマンドの書式は以下の通りです。

```bash

$ sacloud-upload-image [オプション] [イメージ名]

```

#### オプション

    -token  : さくらのクラウドのAPIキー(アクセストークン)
    -secret : さくらのクラウドのAPIキー(シークレット)
    -zone   : 作成するゾーン (is1a/is1b/tk1a) デフォルト:is1a
    -file   : アップロードするISOファイルのパス

`zone`は以下の値をとります
>
    is1a : 石狩第１ゾーン
    is1b : 石狩第２ゾーン
    tk1a : 東京第１ゾーン

#### APIキーの設定方法

上記のオプションで指定する、または環境変数も利用できます。

* オプションで指定する場合

```オプションで指定する場合
$ sacloud-upload-image -token=[アクセストークン] -secret=[シークレット] [イメージ名]
```

* 環境変数で指定する場合

```環境変数で指定する場合
$ export SAKURACLOUD_ACCESS_TOKEN=[アクセストークン]
$ export SAKURACLOUD_ACCESS_TOKEN_SECRET=[シークレット]
$ sacloud-upload-image [イメージ名]
```

#### ISOファイルの指定方法

ISOファイルは`-file`オプション、またはパイプ/リダイレクトで指定します。
`curl`などでISOファイルを取得、パイプで渡してさくらのクラウドへアップロードという使い方ができます。

```ISOファイル指定方法の例

# -fileオプションで指定(CoreOS.isoというファイルをアップロードする例)
$ sacloud-upload-image [オプション] -file ./CoreOS.iso [イメージ名] 

# curlからパイプで受け取る
$ curl -L http://[ISOイメージのURL] | sacloud-upload-image [オプション] [イメージ名]

# リダイレクトで受け取る(CoreOS.isoというファイルをアップロードする例)
$ sacloud-upload-image [オプション] [イメージ名] < CoreOS.iso

```


### 利用例：curlでweb上からISOイメージをダウンロード & さくらのクラウドへアップロード

#### CentOS Atomic Hostの例

```bash
# APIキーを環境変数で指定する
$ export SAKURACLOUD_ACCESS_TOKEN=[アクセストークン]
$ export SAKURACLOUD_ACCESS_TOKEN_SECRET=[シークレット]

# CentOS Atomic Host(http://www.projectatomic.io)の例
$ curl -L http://cloud.centos.org/centos/7/atomic/images/CentOS-Atomic-Host-7-Installer.iso | \
  ./sacloud-upload-image "CentOS Atomic Host"

```

#### CoreOSの例

```bash
# APIキーを環境変数で指定する
$ export SAKURACLOUD_ACCESS_TOKEN=[アクセストークン]
$ export SAKURACLOUD_ACCESS_TOKEN_SECRET=[シークレット]

# CoreOS(https://coreos.com)の例
$ curl -L https://stable.release.core-os.net/amd64-usr/current/coreos_production_iso_image.iso | \
  ./sacloud-upload-image "CoreOS stable"

```


## License

This project is published under [Apache 2.0 License](LICENSE).

## Author

* Kazumichi Yamamoto ([@yamamoto-febc](https://github.com/yamamoto-febc))
