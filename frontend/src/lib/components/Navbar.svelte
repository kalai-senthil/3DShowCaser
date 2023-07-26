<script>
	import NavItem from './NavItem.svelte';

	const links = [
		{
			name: 'Contributers',
			route: '/contributers'
		},
		{
			name: 'About',
			route: '/about'
		}
	];
</script>

<nav class="px-10 py-4 bg-secondary text-white shadow-lg">
	<ul class="flex items-center justify-between gap-x-3">
		<li class="grow"><a href="/">Home</a></li>
		{#each links as link (link.route)}
			<NavItem name={link.name} route={link.route} />
		{/each}
	</ul>
</nav>
