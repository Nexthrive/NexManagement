<script setup>
// import { ref } from 'vue'
import Navbar from "../utils/Navbar.vue";
</script>

<template>
  <div class="container">
    <Navbar />

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
              <button>Edit</button>
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
}

th,
td {
  border: 1px solid #dddddd;
  text-align: left;
  padding: 8px;
  color: var(--brand-dark);
}

th {
  background-color: #f2f2f2;
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
}
</style>


