<script lang="ts">
	import { onMount } from 'svelte';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	export let data: Work | null;
	import Icon from '@iconify/svelte';
	export let token: string;
	import {
		AmbientLight,
		BoxGeometry,
		Color,
		CubeTextureLoader,
		DoubleSide,
		LoadingManager,
		Mesh,
		MeshStandardMaterial,
		PerspectiveCamera,
		Scene,
		WebGLRenderer
	} from 'three';
	import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader';
	import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls';
	import type { Work } from '$lib/types';
	let canvas: HTMLCanvasElement;
	let lightIntensity = 1;
	const camera = new PerspectiveCamera(45, window.innerWidth / window.innerHeight, 0.1, 1000);
	const scene = new Scene();
	const loader = new GLTFLoader(new LoadingManager());
	const light = new AmbientLight(new Color('white'), lightIntensity);
	let showSidePanel = false;
	const textureLoader = new CubeTextureLoader();
	textureLoader.setPath('/textures/Bridge/');
	let textureCube = textureLoader.load([
		'px.jpg',
		'nx.jpg',
		'py.jpg',
		'ny.jpg',
		'pz.jpg',
		'nz.jpg'
	]);
	let backgroundIntensity = 1;
	onMount(() => {
		scene.background = textureCube;
		scene.backgroundIntensity = backgroundIntensity;
		loader.setWithCredentials(true);
		loader.setPath('');
		loader.load(`${PUBLIC_BACKEND_URL}/${data?.path}`, (gltf) => {
			gltf.scene.traverse(function (node) {
				if (node instanceof Mesh) {
					node.castShadow = true;
					node.material.side = DoubleSide;
				}
			});
			scene.add(gltf.scene);
		});
		const renderer = new WebGLRenderer({
			canvas
		});
		const controls = new OrbitControls(camera, renderer.domElement);
		controls.enableDamping = true;
		controls.enableZoom = true;
		function render() {
			window.requestAnimationFrame(() => {
				light.intensity = lightIntensity;
				scene.backgroundIntensity = backgroundIntensity;
				controls.update();
				camera.updateProjectionMatrix();
				renderer.render(scene, camera);
				window.requestAnimationFrame(render);
			});
		}
		renderer.setSize(window.innerWidth, window.innerHeight);
		scene.add(light);
		camera.position.z = 5;
		scene.add(camera);
		render();
		window.addEventListener('resize', () => {
			renderer.setSize(window.innerWidth, window.innerHeight);
			controls.update();
			camera.updateProjectionMatrix();
			camera.aspect = window.innerWidth / window.innerHeight;
			renderer.render(scene, camera);
		});
	});
</script>

<section class="relative">
	<div
		class={`absolute flex top-20 transition-all ease-out items-start ${
			showSidePanel ? 'right-0' : '-right-52'
		}`}
	>
		<button
			on:click={() => {
				showSidePanel = !showSidePanel;
			}}
			class="bg-black rounded-l-md p-1 cursor-pointer mt-10"
		>
			<Icon
				icon="bxs:up-arrow"
				class={`duration-500 ${
					showSidePanel ? 'rotate-90' : '-rotate-90'
				} transition-transform text-2xl ease-in-out`}
			/>
		</button>
		<section class="bg-black w-52 shadow-lg p-2 rounded-md">
			<div class="flex flex-col">
				<div class="flex">
					<h4 class="grow">Light Intensity</h4>
					<p class="border p-1 aspect-square align-text-bottom w-10 text-center">
						{lightIntensity}
					</p>
				</div>
				<input min="0" max="100" step="1" type="range" bind:value={lightIntensity} />
			</div>
			<div class="flex flex-col">
				<div class="flex">
					<h4 class="grow">Background Intensity</h4>
					<p class="border p-1 aspect-square align-text-bottom w-10 text-center">
						{backgroundIntensity}
					</p>
				</div>
				<input min="0" max="10" step=".5" type="range" bind:value={backgroundIntensity} />
			</div>
			<div class="flex flex-col">
				<h4 class="grow">Environment</h4>
				<select
					on:change={(e) => {
						textureLoader.setPath(e.currentTarget.value);
						textureCube = textureLoader.load([
							'px.jpg',
							'nx.jpg',
							'py.jpg',
							'ny.jpg',
							'pz.jpg',
							'nz.jpg'
						]);
						scene.background = textureCube;
					}}
				>
					<option value="/textures/Bridge/">Bridge</option>
					<option value="/textures/Yokahama/">Yokahama</option>
					<option value="/textures/Storfrozen/">Storfrozen</option>
				</select>
			</div>
		</section>
	</div>
	<canvas bind:this={canvas} />
</section>
