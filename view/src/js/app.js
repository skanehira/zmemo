import Vue from 'vue';
import VueRouter from 'vue-router';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import locale from 'element-ui/lib/locale/lang/ja'
import routes from './routes';
import axios from 'axios';
import store from '../store/index';

Vue.use(VueRouter);
Vue.use(ElementUI, { locale })

// コンポーネント
import '../css/app.css';
import menu from '../components/menu/menu.vue';
// アイコンセット
import "ionicons/dist/css/ionicons.css";

// 全コンポーネントでaxiosを使用できる様にprototypeに登録
Vue.prototype.$axios = axios.create({
    // baseURL: '/Zmemo/',
    headers: {
        'ContentType': 'application/json',
        'X-Requested-With': 'XMLHttpRequest'
    },
    responseType: 'json'
});

const router = new VueRouter({
    routes: routes
});

const app = new Vue({
    el: '#app',
    components: {
        'header-menu': menu,
    },
    router,
    store
});

// 画面遷移時の処理
app.$router.beforeEach((to, from, next) => {
    // チェック対象外パス
    let ignorePaths = app.$store.getters.ignorePaths;
    let isNext = false;

    // ログインチェック対象外画面
    for (let path of ignorePaths) {
        if (to.path === path) {
            isNext = true;
            continue;
        }
    }

    if (app.$store.getters.userInfo.userName === "" && !isNext) {
        next("/login");
    } else {
        next();
    }
});