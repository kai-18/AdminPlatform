<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title>
          Quasar App
        </q-toolbar-title>

        <q-space />
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above bordered>
      <q-list>
        <q-item-label header>Admin
          <q-btn-toggle style="float: right;"
            v-model="selectedLanguage"
            toggle-color="light-blue"
            :options="[
              {label: '🇬🇧 EN', value: 'en'},
              {label: '🇮🇹 IT', value: 'it'}
            ]"
            dense
            unelevated
          />

        </q-item-label>

        <EssentialLink
          v-for="link in currentLinks"
          :key="link.title"
          v-bind="link"
        />
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import EssentialLink from 'components/EssentialLink.vue'
import api from '/api/api.js'

const leftDrawerOpen = ref(false)
const selectedLanguage = ref('en') // Default language

const linksListEN = [
  { title: 'Dashboard',
     caption: 'Overview of sales, employees, and reports',
     icon: 'dashboard',
    link: ''
  },
  {  title: 'Employees',
     caption: 'Add, edit, and manage employee details',
     icon: 'badge',
    link: ''
  },
  {  title: 'Sales',
     caption: 'View and track sales records',
     icon: 'attach_money',
    link: ''
  },
  {  title: 'Reports & Analytics',
     caption: 'Generate performance reports',
     icon: 'bar_chart',
    link: ''
  },
  {  title: 'User Management',
     caption: 'Manage admin & staff roles (if needed)',
     icon: 'supervisor_account',
    link: ''
  },
  {  title: 'Settings',
     caption: 'Configure system preferences',
     icon: 'settings',
    link: ''
  },
  {  title: 'Logout',
     caption: '',
     icon: 'logout',
    link: '' }
]

const linksListIT = [
  {
    title: 'Dashboard',
     caption: 'Panoramica di vendite, dipendenti e report',
     icon: 'dashboard',
     link: ''
  },
  {
    title: 'Dipendenti',
     caption: 'Aggiungi, modifica e gestisci i dettagli dei dipendenti',
     icon: 'badge',
     link: ''
  },
  {
    title: 'Vendite',
     caption: 'Visualizza e monitora i record delle vendite',
     icon: 'attach_money',
     link: ''
  },
  {
    title: 'Report e Analisi',
     caption: 'Genera report sulle prestazioni',
     icon: 'bar_chart',
     link: ''
  },
  {
    title: 'Gestione Utenti',
     caption: 'Gestisci ruoli di amministratori e personale (se necessario)',
     icon: 'supervisor_account',
     link: ''
  },
  {
    title: 'Impostazioni',
     caption: 'Configura le preferenze di sistema',
     icon: 'settings',
     link: ''
  },
  {
    title: 'Esci',
     caption: '',
     icon: 'logout',
     link: ''
  }
]

// Computed property to switch menu based on selected language
const currentLinks = computed(() => selectedLanguage.value === 'en' ? linksListEN : linksListIT)

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value
}

const user = ref({})

const fetchMockData = async () => {
  try {
    const response = await api.getUser()
    user.value = response.data.user
  } catch (error) {
    console.error('Error fetching mock data:', error)
  }
}

onMounted(() => {
  fetchMockData()
})
</script>
