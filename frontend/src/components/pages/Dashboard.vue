<script setup>
import { ref } from 'vue'
import Navbar from "../utils/Navbar.vue";
</script>

<template>
  <div class="container">
    <Navbar />
    <!-- dashboard content -->
    <div class="dashboard-container">
      <!-- header -->
      <div class="header">
        <h1>Welcomez</h1>
        <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Quasi exercitationem aut qui ratione <br> necessitatibus
          id sunt eveniet, blanditiis laudantium cupiditate vel quaerat at, tempora, iure iusto eos!</p>
      </div>

      <!-- dashboard-content -->
      <div class="card-container">
        <!-- list card -->

        <div class="card" >
            <h2>Jumlah Barang</h2>
            <h1>770</h1>
        </div>

        <div class="card">
            <h2>Stock Total</h2>
            <h1>100</h1>
        </div>

      </div>
    </div>
  </div>
</template>

<style scoped>
.dashboard-container {
  width: 100vw;
  margin: 20px;
}
.container {
  width: 100vw;
  height: 100vh;
}

.header {
  font-size: 20px;
  color: var(--brand-dark);
  margin: 10px;
  padding-top: 20px;
}

.card-container {
  color:var(--brand-dark);
  margin-top: 50px;
  display: flex;
  gap: 20px;
}

.card {
  background-color: var(--brand-primary);
  width: 250px;
  padding: 10px;
  color: var(--brand-light);
  border-radius: 10px;
}

.card h1 {
  font-size: 50px;
  margin: 0;
  padding: 0;
}
.card h2 {
  margin: 0;
  padding: 0;
}

</style>
