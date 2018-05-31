<template src="./memoList.html"></template>
<style scoped src="./memoList.css"></style>


<script>
export default {
  data() {
    return {
      userName: "",
      memoList: []
    };
  },
  methods: {
    async getMemoList() {
      await this.$axios
        .get("/memos/" + this.userName)
        .then(response => {
          this.memoList = response.data;
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
    async deleteMemo(memo) {
      await this.$axios
        .delete("/memos/" + memo.userName + "/" + memo.memoId)
        .then(response => {
          this.$alert("削除しました", "", {
            confirmButtonText: "はい",
            type: "info",
            callback: () => {
              this.getMemoList();
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
    showMemo() {
      // todo show memo in dialog
      console.log("show memo");
    },
    dragstart(item, e) {
      console.log(item);
      console.log(e);
    },
    dragend(e) {
      console.log(e);
    },
    dragenter(item) {
      console.log(item);
    }
  },
  mounted() {
    this.userName = this.$store.getters.userInfo.userName;
    this.getMemoList();
  }
};
</script>

