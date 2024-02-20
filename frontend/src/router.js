import { createRouter, createWebHistory } from "vue-router";


import Dashboard from "./components/pages/Dashboard.vue";
import UserManage from "./components/pages/UserManage.vue";
import Inventory from "./components/pages/Inventory.vue";

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
		path: "/usermanage",
		component: UserManage,
		name: "UserManage",
		// meta: {
		// 	requiresAuth: true,
		// },
	},
	{
		path: "/inventory",
		component: Inventory,
		name: "Inventory",
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

