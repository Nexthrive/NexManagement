<script setup>
	// import { ref } from 'vue'
	import Navbar from "../utils/Navbar.vue";
</script>

<template>
	<div class="container">
		<Navbar />
		<!-- Welcome section -->
		<div class="dashboard-container">
			<div class="header">
				<h1>Welcome to <span class="highlight">NexManagement</span></h1>
				<p>
					Unlock the future of seamless business operations with NexManagement!
					Immerse yourself in a world of innovation, where
					<span class="highlight">efficiency</span> meets
					<span class="highlight">intuition</span>. Lorem ipsum dolor sit amet,
					consectetur adipisicing elit. Quasi exercitationem aut qui ratione
					necessitatibus id sunt eveniet, blanditiis laudantium cupiditate vel
					quaerat at, tempora, iure iusto eos!
				</p>
			</div>

			<!-- Dashboard content -->
			<div class="card-container">
				<!-- Statistics cards -->
				<div class="card">
					<h2>Total Products</h2>
					<h1>770</h1>
				</div>

				<div class="card">
					<h2>Total Stock</h2>
					<h1>100</h1>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
	.container {
		width: 80vw;
		height: 100vh;
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		margin-top: 30px;
	}

	.dashboard-container {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}

	.header {
		font-size: 20px;
		color: var(--brand-dark);
		margin: 10px;
		padding-top: 20px;
	}

	.header {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}

	.header p {
		text-align: center;
	}

	.card-container {
		color: var(--brand-dark);
		margin-top: 10px;
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
