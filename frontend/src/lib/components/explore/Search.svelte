<script>
	import TagItems from './TagItems.svelte';
	import Model from './Model.svelte';
	const items = [
		{
			name: 'Objects',
			route: '/'
		},
		{
			name: 'Vehivles',
			route: '/'
		},
		{
			name: 'Buildings',
			route: '/'
		},
		{
			name: 'Scenary',
			route: '/'
		},
		{
			name: 'Tree',
			route: '/'
		},
		{
			name: 'SuperHero',
			route: '/'
		},
		{
			name: 'Characters',
			route: '/'
		},
		{
			name: 'AirCrafts',
			route: '/'
		},
		{
			name: 'Weapons',
			route: '/'
		},
		{
			name: 'Animals',
			route: '/'
		},
		{
			name: 'Sports',
			route: '/'
		},
		{
			name: 'Plants',
			route: '/'
		}
	];
</script>

<div>
	<ul class="flex justify-evenly py-3">
		{#each items as item}
			<TagItems title={item.name} />
		{/each}
	</ul>
	<form class="flex justify-end py-2 pr-10">
		<input type="text" color="pink" placeholder="&nbsp;&nbsp;Search 3D models ..." value="" />
	</form>
	{#each items as item}
		<Model name={item.name} route={item.route} />
	{/each}
</div>
