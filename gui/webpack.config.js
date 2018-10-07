module.exports = {
        // ビルド対象js
        entry: './src/js/app.js',
        output: {
                // ビルド後出力するファイル
                path: __dirname,
                filename: '../public/bundle.js'
        },
        resolve: {
                // webpackでvueをビルドできるようにするためのおまじない
                alias: {
                        vue: "vue/dist/vue.js"
                }
        },
        module: {
                // ファイルタイプ毎にローダーを設定
                // testに対象ファイル拡張子（正規表現可能）を指定
                rules: [
                        {
                                test: /\.vue$/,
                                loader: 'vue-loader',
                        },
                        {
                                test: /\.css$/,
                                loader: 'style-loader!css-loader', // css
                        },
                        {
                                test: /\.(otf|eot|svg|ttf|woff|woff2)(\?.+)?$/, // フォントやアイコン
                                loader: 'url-loader',
                        },
                        {
                                test: /\.(jpg|png|gif)$/, // imageファイル
                                loaders: 'url-loader'
                        },
                ]
        },
};
