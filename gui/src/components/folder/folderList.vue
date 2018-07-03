<template src="./folderList.html"></template>
<style scoped src="./folderList.css"></style>


<script>
export default {
  data() {
    return {
      userName: "",
      folderList: []
    };
  },
  methods: {
    async getFolderList() {
      await this.$axios
        .get("/folders/" + this.userName)
        .then(response => {
          this.folderList = response.data;
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
    async deleteFolder(folder) {
      await this.$axios
        .delete("/folders/" + folder.userName + "/" + folder.folderName)
        .then(response => {
          this.$alert("削除しました", "", {
            confirmButtonText: "はい",
            type: "info",
            callback: () => {
              this.getFolderList();
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
    }
  },
  mounted() {
    this.userName = this.$store.getters.userInfo.userName;
    this.getFolderList();
  }
};
</script>

