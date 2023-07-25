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

<nav class="p-2 bg-black text-white shadow-md">
	<ul class="flex items-center justify-between gap-x-3">
		<li class="grow"><a href="/">Home</a></li>
		{#each links as link (link.route)}
			<NavItem name={link.name} route={link.route} />
		{/each}
	</ul>
</nav>
