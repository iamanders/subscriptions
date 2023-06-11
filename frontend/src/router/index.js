import { createRouter, createWebHistory, createWebHashHistory } from 'vue-router';
import Home from '../views/Home.vue';

const routes = [
  { path: '/', name: 'Home', component: Home },
  { path: '/settings', name: 'Settings', component: () => import(/* webpackChunkName: "settings" */ '../views/Settings.vue') },
  { path: '/stats', name: 'Stats', component: () => import(/* webpackChunkName: "stats" */ '../views/Stats.vue') },
  { path: '/subscription/new', name: 'NewSubscription', component: () => import(/* webpackChunkName: "new-subscription" */ '../views/NewSubscription.vue') },
  { path: '/subscription/:id', name: 'EditSubscription', component: () => import(/* webpackChunkName: "edit-subscription" */ '../views/EditSubscription.vue') },
  { path: '/categories', name: 'CategoriesList', component: () => import(/* webpackChunkName: "list-categories" */ '../views/CategoriesList.vue') },
  { path: '/subscription/new', name: 'NewCategory', component: () => import(/* webpackChunkName: "new-category" */ '../views/NewCategory.vue') },
  { path: '/categories/:id', name: 'EditCategory', component: () => import(/* webpackChunkName: "edit-category" */ '../views/EditCategory.vue') },
];

const router = createRouter({
  history: process.env.VUE_APP_HASH_MODE === 'true' ? createWebHashHistory() : createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
