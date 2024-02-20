<script setup>
// import { ref } from 'vue'
import Navbar from "../utils/Navbar.vue";
import EditForm from "../utils/EditForm.vue";
import AddForm from "../utils/AddForm.vue";
import { ref } from "vue";

const showEditForm = ref(false);
const showEditFormModal = () => {
  showEditForm.value = true;
};
const closeFormModal = () => {
  showEditForm.value = false;
};


const showAddForm = ref(false);
const showAddFormModal = () => {
  showAddForm.value = true;
};
const closeAddModal = () => {
  showAddForm.value = false;
};

</script>

<template>
  <div v-if="showAddForm" class="overlay"></div>
  <div v-if="showEditForm" class="overlay"></div>
  <EditForm v-if="showEditForm" @close-form-modal="closeFormModal" :class="{ 'fade-scale-in': showEditForm }" />
  <AddForm v-if="showAddForm" @close-add-modal="closeAddModal" :class="{ 'fade-scale-in': showAddForm }" />
  <div class="container">
    <Navbar />
    <button class="add-button" @click="showAddFormModal">Add User</button>
    <div class="data-user-container">
      <table>
        <thead>
          <tr>
            <th>Username</th>
            <th>Email</th>
            <th>No Telepon</th>
            <th>Role</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>Nithze</td>
            <td>Nithze@gmail.com</td>
            <td>123456789</td>
            <td>User</td>
            <td class="action-button">
              <button @click="showEditFormModal">Edit</button>
              <button>Delete</button>
            </td>
          </tr>
          <tr>
            <td>Nithze</td>
            <td>Nithze@gmail.com</td>
            <td>123456789</td>
            <td>User</td>
            <td class="action-button">
              <button>Edit</button>
              <button>Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
.container {
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.data-user-container {
  margin-top: 40px;
  width: 90vw;
  height: 100vh;
  align-self: center;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
  overflow: hidden;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

th,
td {
  border: none;
  /* text-align: left; */
  padding: 8px;
  color: var(--brand-dark);
}


th {
  background-color: #f2f2f2;
}

tbody tr:nth-child(even) {
  background-color: #f9f9f9;
}

.action-button {
  padding: 8px;
  text-align: center;
  display: flex;
  justify-content: space-around;
}

.action-button button {
  background-color: var(--brand-primary);
  color: var(--brand-light);
  border: none;
  padding: 5px 10px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 12px;
  margin: 2px 2px;
  cursor: pointer;
  border-radius: 3px;
  transition: background-color 0.3s ease;
}

.action-button button:hover {
  background-color: var(--brand-primary-hover);
}

.add-button {
  width: 100px;
  margin-left: 87%;
  background-color: var(--brand-primary);
  border: none;
  border-radius: 3px;
  padding: 5px 10px
}


.overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  z-index: 100;
}


.fade-scale-in {
  animation: fadeScaleIn 0.4s;
}

@keyframes fadeScaleIn {
  from {
    opacity: 0;
    transform: scale(0.8);
  }

  to {
    opacity: 1;
    transform: scale(1);
  }
}

</style>


