import Vue from 'vue';
import Vuex from 'vuex';
import createPersistedState from 'vuex-persistedstate';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    userInfo: {
      userName: "",
    },
    ignorePaths: [
      "/",
      "/login",
      "/logout",
      "/registUser"
    ]
  },
  getters: {
    userInfo(state) {
      return state.userInfo;
    },
    ignorePaths(state) {
      return state.ignorePaths;
    }
  },
  mutations: {
    setLoginUserInfo(state, userInfo) {
      state.userInfo.userName = userInfo.userName;
    },
    logout(state) {
      state.userInfo.userName = "";
    }
  },
  // vuxの情報をセッションに保存する
  plugins: [
    createPersistedState({
      key: 'zmemo',
      storage: window.sessionStorage,
    }),
  ],
});