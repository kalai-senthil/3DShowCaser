<script lang="ts">
	import addIcon from '$lib/assets/add.svg';
</script>

<section
	class="col-span-12 md:col-span-3 rounded-md border-[.5px] border-gray-500 py-5 cursor-pointer hover:bg-gray-700/50 transition-colors"
>
	<a class="justify-center flex-col items-center flex" href="/add">
		<img src={addIcon} alt="add" class="h-20" />
		<h4 class="font-bold text-sm">Add Work</h4>
	</a>
</section>
