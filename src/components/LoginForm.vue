<template>
	<div class="login">
		<form>
			<label>
				Username:
				<input type="text" placeholder="Username" />
			</label>
			<label>
				Password:
				<input type="password" placeholder="Password" />
			</label>
			<input type="submit" />	
		</form>
	</div>
</template>

<script>
export default {
	name: 'LoginForm',
	data () {
		return {
			username: '',
			password: ''
		}
	}
}
</script>

<style>
label {
	float: left;
	display: inline-block;
}
</style>