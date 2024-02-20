<script setup>
	import { ref } from "vue";
	import { useRouter } from "vue-router";
	import { jwtDecode } from "jwt-decode";
	import axios from "axios";

	// primevue
	import InputText from "primevue/inputtext";
	import Password from "primevue/password";

	const router = useRouter();

	// refs
	const Username = ref("");
	const Passphrase = ref("");

	const login = async () => {
		const res = await axios.post("http://localhost:8080/nex/login", {
			username: Username.value,
			passphrase: Passphrase.value,
		});

		const token = res.data.token;

		if (token) {
			const decoded = jwtDecode(token);
			const UserID = decoded.id;

			localStorage.setItem("jwtToken", token);
			localStorage.setItem("UserID", UserID);

			router.push("/");
		}
	};
</script>

<template>
	<div class="container-floating">
		<form @submit.prevent="login">
			<div class="form-header">
				<h3>Log-In!</h3>
			</div>
			<div class="inputs">
				<div class="input">
					<InputText
						id="username"
						v-model="Username"
						placeholder="Username"
						class="p-fluid"
					/>
				</div>
				<div class="input">
					<Password
						id="passphrase"
						v-model="Passphrase"
						:feedback="false"
						placeholder="Passphrase"
						class="p-fluid"
					/>
				</div>
				<button type="submit">Log in</button>
			</div>
		</form>
	</div>
</template>

<style scoped>
	.container-floating {
		width: 100vw;
		height: 100vh;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	body {
		background: var(--brand-gradient);
	}

	form {
		background-color: var(--brand-primary);
		padding: 50px 20px;
		border-radius: 7px;
		width: 30%;
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.form-header h3 {
		color: white;
		font-size: 3rem;
		margin: 0;
		padding: 0;
		margin-bottom: 20px;
	}

	.inputs {
		width: 55%;
	}

	.input {
		margin-bottom: 10px;
		width: 100%;
	}

	:deep(.p-fluid) {
		width: 100%;
	}

	button {
		background-color: white;
		border: none;
		padding: 10px;
		border-radius: 7px;
		width: 100%;
		cursor: pointer;
	}
</style>
