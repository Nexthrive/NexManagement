import { createRouter, createWebHistory } from "vue-router";

import Dashboard from "./components/pages/Dashboard.vue";
import Login from "./components/pages/Login.vue";

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
	},
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});

export default router;
