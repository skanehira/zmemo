<template src="./newFolder.html"></template>
<style scoped src="./newFolder.css"></style>

<script>
export default {
  data() {
    return {
      labelwidth: "100px",
      isVisible: true,
      folder: {
        userName: "",
        folderName: ""
      }
    };
  },
  methods: {
    close() {
      // 差分があった場合は確認ダイアログ
      if (this.isFormHasDiff()) {
        this.$confirm("フォルダ作成を中止しますか？", "Confirm", {
          confirmButtonText: "はい",
          cancelButtonText: "いいえ",
          type: "info"
        })
          .then(() => {
            // 破棄時はトップに戻る
            this.$router.push("/");
          })
          .catch(() => {
            // キャンセル時は何もしない
          });
      } else {
        this.$router.push("/");
      }
    },
    // フォルダ作成
    async createFolder() {
      await this.$axios
        .post("/folders", this.folder)
        .then(response => {
          this.$alert("登録しました", "Info", {
            confirmButtonText: "はい",
            callback: () => {
              this.$router.push("/folderList");
            }
          });
        })
        .catch(error => {
          // ネットワークエラー時
          if (
            error.hasOwnProperty("message") &&
            !error.hasOwnProperty("response")
          ) {
            this.$alert(error.message, "Error", {
              confirmButtonText: "はい",
              type: "error"
            });
            // サーバーエラー
          } else if (error.response.data.hasOwnProperty("message")) {
            this.$alert(error.response.data.message, "Error", {
              confirmButtonText: "はい",
              type: "error"
            });
          }
        });
    },
    // formが更新された場合はtrueを返す
    isFormHasDiff() {
      for (let key of Object.keys(this.folder)) {
        if (key !== "userName" && this.folder[key] !== "") {
          return true;
        }
      }
    }
  },
  mounted() {
    // ユーザ名を取得
    this.folder.userName = this.$store.getters.userInfo.userName;
  }
};
</script>
