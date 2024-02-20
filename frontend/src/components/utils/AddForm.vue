<script setup>
const emits = defineEmits(["closeFormModal"]);
import "primeicons/primeicons.css";
import "primeicons/primeicons.css";
const closeAddModal = () => {
  emits("closeAddModal");
};

</script>

<template>
  <div class="add-form-overlay">
    <div class="add-form-container">
      <!-- header ini aa -->
      <div class="form-header">
        <h1>create user</h1>
        <i class="pi pi-times" @click="closeAddModal"></i>
      </div>

      <!-- form ini aa -->
      <form class="form-container" @submit.prevent="addtask">
        <!-- nama div -->
        <div class="name-input">
          <label for="username">Username</label>
          <input id="username" type="text" v-model="title" />
        </div>


        <div class="name-input">
          <label for="email">Email</label>
          <input id="email" type="email" v-model="title" />
        </div>

        <div class="name-input">
          <label for="telp">No Telefon</label>
          <input id="telp" type="number" v-model="title" />
        </div>

        <div class="name-input">
          <label for="passphrase">passphrase</label>
          <input id="passphrase" type="password" v-model="title" />
        </div>

        <div class="button-form">
          <button class="button-cencel" @click="closeformmodal">cencel</button>
          <button class="button-save" type="submit">save</button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
	.button-form {
		display: flex;
		justify-content: flex-end;
		width: 100%;
		padding: 1.5rem;
		gap: 1rem;
	}

	.button-form .button-save {
		width: 20%;
		background: #f67000;
		border: none;
		padding: 10px 5px;
		border-radius: 5px;
		transition: all 0.2s;
	}

	.button-form .button-save:hover {
		background: #6430ab;
	}

	.button-form .button-cencel {
		width: 20%;
		background: #0a0b0a;
		border: 1px solid #f67000;
		padding: 10px 5px;
		border-radius: 5px;
		transition: all 0.2s;
	}

	.button-form .button-cencel:hover {
		border: 1px solid #ff0000;
	}

	.start-due {
		display: flex;
		width: 100%;
		justify-content: flex-start;
		align-items: center;
		gap: 1.5rem;
		padding-top: 1.4rem;
	}

	.start {
		width: 50%;
		display: flex;
		flex-direction: column;
		gap: 7px;
	}

	.due {
		width: 50%;
		display: flex;
		flex-direction: column;
		gap: 7px;
	}

	.form-container {
		width: 100%;
		padding: 0rem 1.5rem;
		display: flex;
		flex-direction: column;
		gap: 5px;
	}

	.form-header {
		display: flex;
		align-items: center;
		justify-content: space-between;
		flex-shrink: 0;
		width: 100%;
		height: 70px;
		padding: 1.5rem;
	}

	.form-header i {
		cursor: pointer;
	}

	.description-input {
		width: 100%;
		display: flex;
		flex-direction: column;
		gap: 7px;
		margin-top: 10px;
	}

	.name-input {
		width: 100%;
		display: flex;
		flex-direction: column;
		gap: 7px;
	}

	.form-header h1 {
		font-size: 1.4rem;
	}

	.add-form-overlay {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		/* background: rgba(0, 0, 0, 0.5); */
		display: flex;
		justify-content: center;
		align-items: center;
		z-index: 100;
	}

	.add-form-container {
		background: #fff;
		/* padding: 10px; */
		border-radius: 8px;
		box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
		width: 30%;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
    color: #000;
	}

	/* prime-vue-custom */
	label {
		font-size: 0.8rem;
		font-weight: 500;
		color: #000;
	}

	.p-inputtext:enabled {
		border-color: #52525b;
		width: 500%;
	}

	.p-inputtext:enabled:hover {
		border-color: #52525b;
		width: 100%;
	}
  input {
    background-color: #fff;
    border: none;
    border: 1px solid #999999;
    border-radius: 4px;
    padding: 7px;
  }

</style>


