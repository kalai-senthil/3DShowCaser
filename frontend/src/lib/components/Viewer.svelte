<script lang="ts">
	import { onMount } from 'svelte';

	export let data: Work | null;
	import {
		AmbientLight,
		BoxGeometry,
		Color,
		Mesh,
		MeshBasicMaterial,
		MeshStandardMaterial,
		PerspectiveCamera,
		PlaneGeometry,
		Scene,
		WebGLRenderer
	} from 'three';
	import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader';
	import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls';
	import type { Work } from '$lib/types';
	let canvas: HTMLCanvasElement;
	const camera = new PerspectiveCamera(45, window.innerWidth / window.innerHeight, 0.1, 1000);
	const scene = new Scene();
	const loader = new GLTFLoader();
	const material = new MeshBasicMaterial({
		color: 'green'
	});
	const light = new AmbientLight(new Color(0x0000ff));
	const geometry = new BoxGeometry(1, 1, 1);
	const mesh = new Mesh(geometry, material);
	onMount(() => {
		console.log(data);

		// loader.parse(`data:application/octet-stream;base64,${data?.file}`, '', (gltf) => {
		// 	console.log(gltf);
		// });
		const renderer = new WebGLRenderer({
			canvas
		});
		const controls = new OrbitControls(camera, renderer.domElement);
		function render() {
			window.requestAnimationFrame(() => {
				controls.update();
				camera.updateProjectionMatrix();
				camera.updateMatrix();
				renderer.render(scene, camera);
				window.requestAnimationFrame(render);
			});
		}
		renderer.setSize(window.innerWidth, window.innerHeight);
		scene.add(mesh);
		scene.add(light);
		camera.position.z = 5;
		scene.add(camera);
		render();
		window.addEventListener('resize', () => {
			controls.update();
			camera.updateProjectionMatrix();
			renderer.setSize(window.innerWidth, window.innerHeight);
			camera.updateProjectionMatrix();
			renderer.render(scene, camera);
		});
	});
</script>

<canvas bind:this={canvas} />
