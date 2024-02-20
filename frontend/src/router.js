import { createRouter, createWebHistory } from "vue-router";


import Dashboard from "./components/pages/Dashboard.vue";

const routes = [
	{
		path: "/",
		component: Dashboard,
		name: "Dashboard",
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

