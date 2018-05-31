<template src="./login.html"></template>
<style scoped src="./login.css"></style>

<script>
export default {
  data() {
    return {
      labelWidht: "70px",
      user: {
        userName: "",
        password: ""
      }
    };
  },
  methods: {
    async login() {
      await this.$axios
        .post("/users/login", this.user)
        .then(response => {
          this.$store.commit("setLoginUserInfo", response.data);

          this.$alert("ログインしました", "Info", {
            confirmButtonText: "はい",
            type: "success",
            callback: () => {
              this.$router.push("/");
            }
          });
        })
        .catch(error => {
          // ネットワークエラー時
          if (
            error.hasOwnProperty("message") &&
            !error.hasOwnProperty("response")
          ) {
            this.$alert(error.message, "エラー", {
              confirmButtonText: "はい",
              type: "error"
            });
            // サーバーエラー
          } else if (error.response.data.hasOwnProperty("message")) {
            this.$alert(error.response.data.message, "エラー", {
              confirmButtonText: "はい",
              type: "error"
            });
          }
        });
    },
    registUser() {
      // ユーザ登録画面に遷移
      this.$router.push("/registUser");
    }
  }
};
</script>

