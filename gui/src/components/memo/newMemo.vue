<template src="./newMemo.html"></template>
<style scoped src="./newMemo.css"></style>

<script>
export default {
  data() {
    return {
      labelwidth: "65px",
      isVisible: true,
      memo: {
        userName: "",
        title: "",
        text: ""
      }
    };
  },
  methods: {
    close() {
      // 差分があった場合は確認ダイアログ
      if (this.isFormHasDiff()) {
        this.$confirm("メモを破棄しますか？", "Confirm", {
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
    // メモ作成
    async saveMemo() {
      await this.$axios
        .post("/memos", this.memo)
        .then(response => {
          this.$alert("登録しました", "Info", {
            confirmButtonText: "はい",
            callback: () => {
              this.$router.push("/memoList");
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
    // formが更新された場合はtrueを返す
    isFormHasDiff() {
      for (let key of Object.keys(this.memo)) {
        if (key !== "userName" && this.memo[key] !== "") {
          return true;
        }
      }
    }
  },
  async mounted() {
    // ユーザ名を取得
    this.memo.userName = this.$store.getters.userInfo.userName;
  }
};
</script>
