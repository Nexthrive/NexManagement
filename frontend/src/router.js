import { createRouter, createWebHistory } from "vue-router";

import Dashboard from "./components/pages/Dashboard.vue";
import Login from "./components/pages/Login.vue";
import UserManage from "./components/pages/UserManage.vue";

const routes = [
	{
		path: "/",
		component: Dashboard,
		name: "Dashboard",
		meta: {
			requiresAuth: true,
		},
	},
	{
		path: "/usermanage",
		component: UserManage,
		name: "UserManage",
		meta: {
			requiresAuth: true,
		},
	},
	{
		path: "/login",
		component: Login,
		name: "Login",
	},
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});

router.beforeEach((to, from, next) => {
	const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);

	if (requiresAuth) {
		const jwtToken = localStorage.getItem("jwtToken");

		if (!jwtToken) {
			next({ name: "Login" });
		} else {
			next();
		}
	} else {
		next();
	}
});

export default router;
