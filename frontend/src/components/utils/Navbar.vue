<script setup>
</script>

<template>
  <div class="nav">
    <div class="nav-container">
      <div class="nav-logo">
        <!-- <img src="../assets/logo.png" alt="logo"> -->
        <h3>Nex<span>Management</span>
        </h3>
      </div>
      <div class="nav-links">
        <router-link to="/">Dashboard</router-link>
        <router-link to="/about">Inevtory</router-link>
        <router-link to="/about">About</router-link>
      </div>
    </div>
  </div>
</template>

<style scoped>
.nav-container {
  display: flex;
  justify-content: center;
  align-items: center;
}

.nav-links {
  width: 80%;
  display: flex;
  gap: 2rem;
  justify-content: center;
  margin-right: 20%;
  align-items: center;
}
.nav-logo {
  width: 20%;
  margin-left: 10px;
  align-items: center;
  color: #000;
  font-size: 17px;
}

.nav-links a {
  text-decoration: none;
  color: #000;
}
span {
  color: var(--brand-primary);
}
</style>


