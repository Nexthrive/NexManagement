import { createRouter, createWebHistory } from "vue-router";

import Dashboard from "./components/pages/Dashboard.vue";
import Login from "./components/pages/Login.vue";
import UserManage from "./components/pages/UserManage.vue";

const routes = [
	{
		path: "/",
		component: Dashboard,
		name: "Dashboard",
		// meta: {
		// 	requiresAuth: true,
		// },
	},
	{
		path: "/login",
		component: Login,
		name: "Login",
		path: "/usermanage",
		component: UserManage,
		name: "UserManage",
		// meta: {
		// 	requiresAuth: true,
		// },
	},
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});

export default router;