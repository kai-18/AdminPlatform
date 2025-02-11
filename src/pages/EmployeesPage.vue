<template>
  <div class="q-pa">
    <q-table class="my-sticky-dynamic" flat bordered title="Employees" :rows="rows" :columns="columns"
      :loading="loading" row-key="index" virtual-scroll :virtual-scroll-item-size="48"
      :virtual-scroll-sticky-size-start="48" :pagination="pagination" :rows-per-page-options="[0]" />
    <!-- @virtual-scroll="onScroll" -->
  </div>
</template>

<script>
import { ref, /*nextTick,*/ onMounted } from 'vue'

const columns = [
  {
    name: 'name',
    required: true,
    label: 'Name',
    field: row => row.name,
    align: 'left',
    sortable: true
  },
  {
    name: 'lastname',
    label: 'Lastname',
    field: 'lastname',
    align: 'left',
    sortable: true
  },
  {
    name: 'username',
    label: 'Username',
    field: 'username',
    align: 'left',
    sortable: true
  },
  {
    name: 'email',
    label: 'Email',
    field: 'email',
    align: 'left',
    sortable: true
  },
  {
    name: 'position',
    label: 'Position',
    field: 'position',
    align: 'left',
    sortable: true
  }
]

export default {
  setup() {
    const rows = ref([]) // Store employees
    //    const nextPage = ref(2) // Pagination tracker
    const loading = ref(false)

    const fetchEmployees = async () => {
      loading.value = true
      try {
        const response = await fetch('http://localhost:8080/users')
        if (!response.ok) throw new Error('Failed to fetch employees')

        const data = await response.json()
        rows.value = data.map((row, index) => ({ ...row, index }))
      } catch (error) {
        console.error(error)
      }
      loading.value = false
    }

    onMounted(fetchEmployees) // Fetch data on component mount

    return {
      columns,
      rows,
      loading,
      pagination: { rowsPerPage: 0 },

      // Handles infinite scrolling
      //      onScroll({ to, ref }) {
      //        if (loading.value || to !== rows.value.length - 1) return
      //
      //        loading.value = true
      //        setTimeout(() => {
      //          nextPage.value++
      //          nextTick(() => {
      //            ref.refresh()
      //            loading.value = false
      //          })
      //        }, 500)
      //      }
    }
  }
}
</script>

<style lang="sass">
.my-sticky-dynamic
  height: 700px

  .q-table__top,
  .q-table__bottom,
  thead tr:first-child th
    background-color: #00b4ff

  thead tr th
    position: sticky
    z-index: 1
  thead tr:last-child th
    top: 48px
  thead tr:first-child th
    top: 0
  tbody
    scroll-margin-top: 48px
</style>
