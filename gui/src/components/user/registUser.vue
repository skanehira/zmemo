<template src="./registUser.html"></template>
<style scoped src="./registUser.css"></style>

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
    async regist() {
      // todo confirm dialog
      await this.$confirm("登録しますか？", "Confirm", {
        confirmButtonText: "はい",
        cancleButtonText: "キャンセル",
        type: "info"
      })
        .then(() => {
          this.$axios
            .post("/users", this.user)
            .then(response => {
              this.$store.commit("setLoginUserInfo", response.data);

              this.$alert("登録しました", "Info", {
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
        })
        .catch(() => {
          // キャンセル時は何もしない
        });
    },
    cancel() {
      // topに戻る
      this.$router.push("/");
    }
  }
};
</script>
