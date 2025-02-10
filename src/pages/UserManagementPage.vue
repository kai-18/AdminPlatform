<template>
  <q-page>
    <q-page-container>
      <q-card>
        <q-card-section>
          <q-card-title>
            User Management
          </q-card-title>
        </q-card-section>
        <q-card-section>
          <q-btn
            color="primary"
            label="Create User"
            @click="createUser"
          />
        </q-card-section>
        <q-card-section>
          <q-table
            :rows="users"
            :columns="columns"
            row-key="id"
          >
            <template v-slot:body-cell-actions="props">
              <q-td :props="props">
                <q-btn
                  color="primary"
                  label="Edit"
                  @click="editUser(props.row)"
                />
                <q-btn
                  color="negative"
                  label="Delete"
                  @click="deleteUser(props.row)"
                />
              </q-td>
            </template>
          </q-table>
        </q-card-section>
      </q-card>
    </q-page-container>
  </q-page>
</template>
