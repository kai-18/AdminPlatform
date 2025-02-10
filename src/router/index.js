import { createRouter, createWebHistory } from 'vue-router'
import MainLayout from 'layouts/MainLayout.vue'
import Dashboard from 'pages/DashboardPage.vue'
import Employees from 'pages/EmployeesPage.vue'
import Sales from 'pages/SalesPage.vue'
import ReportsAndAnalytics from 'pages/ReportsAndAnalyticsPage.vue'
import UserManagement from 'pages/UserManagementPage.vue'
import Settings from 'pages/SettingsPage.vue'

const routes = [
  {
    path: '/',
    component: MainLayout,
    children: [
      { path: 'Dashboard', component: Dashboard },
      { path: 'Employees', component: Employees },
      { path: 'Sales', component: Sales },
      { path: 'ReportsAndAnalytics', component: ReportsAndAnalytics },
      { path: 'UserManagement', component: UserManagement },
      { path: 'Settings', component: Settings }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router
