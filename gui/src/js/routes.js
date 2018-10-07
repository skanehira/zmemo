import top from '../components/top/top.vue';
import newMemo from '../components/memo/newMemo.vue';
import memoList from '../components/memo/memoList.vue';
import newFolder from '../components/folder/newFolder.vue';
import folderList from '../components/folder/folderList.vue';
import registUser from '../components/user/registUser.vue';
import login from '../components/user/login.vue';

export default [
        { path: "/", component: top },
        { path: "/newMemo", component: newMemo },
        { path: "/memoList", component: memoList },
        { path: "/newFolder", component: newFolder },
        { path: "/folderList", component: folderList },
        { path: "/registUser", component: registUser },
        { path: "/login", component: login }
]
