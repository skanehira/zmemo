<template src="./menu.html"></template>
<style scoped src="./menu.css"></style>

<script>
export default {
  data() {
    return {
      menus: [],
      loggnedinMenus: [
        {
          name: "メモ",
          icon: "ion-ios-create",
          submenus: [
            { name: "新規作成", path: "newMemo" },
            { name: "メモ一覧", path: "memoList" }
          ]
        },
        {
          name: "フォルダ",
          icon: "ion-ios-folder",
          submenus: [
            { name: "新規作成", path: "newFolder" },
            { name: "フォルダ一覧", path: "folderList" }
          ]
        },
        {
          name: "ユーザ",
          icon: "ion-ios-person",
          submenus: [
            { name: "パスワード変更", path: "changePassword" },
            { name: "ユーザ名変更", path: "changeUserName" },
            { name: "ログアウト", path: "logout" }
          ]
        }
      ],
      notLoggedMenus: [
        {
          name: "ユーザ",
          icon: "ion-ios-person",
          submenus: [
            { name: "ユーザ登録", path: "registUser" },
            { name: "ログイン", path: "login" }
          ]
        }
      ]
    };
  },
  methods: {
    handleSelect(menu) {
      if (menu === undefined) {
        this.$router.push("/");
      } else if (menu === "logout") {
        this.logout();
      } else {
        this.$router.push("/" + menu);
      }
    },
    async logout() {
      await this.$store.commit("logout");

      await this.$alert("ログアウトしました", "Info", {
        submitButtonText: "はい",
        type: "info"
      });

      this.$router.push("/");
    },
    isLoggedin() {
      return this.$store.getters.userInfo.userName === "" ? false : true;
    },
    setMenu() {
      this.menus = this.isLoggedin()
        ? this.loggnedinMenus
        : this.notLoggedMenus;
    }
  },
  mounted() {
    this.setMenu();
  },
  computed: {
    // vuexのstateが変更された場合、コンポーネントにも反映するようにget()を使用
    // https://vuex.vuejs.org/ja/guide/forms.html#双方向算出プロパティ
    userName: {
      get() {
        return this.$store.getters.userInfo.userName;
      }
    }
  },
  // ログイン状態を監視
  watch: {
    userName: function() {
      this.setMenu();
    }
  }
};
</script>

