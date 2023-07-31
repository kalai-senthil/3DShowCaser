<script lang="ts">
	import { onMount } from 'svelte';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	export let data: Work | null;
	export let token: string;
	import {
		AmbientLight,
		BoxGeometry,
		Color,
		DoubleSide,
		LoadingManager,
		Mesh,
		MeshBasicMaterial,
		MeshStandardMaterial,
		PerspectiveCamera,
		Scene,
		Triangle,
		WebGLRenderer
	} from 'three';
	import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader';
	import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls';
	import type { Work } from '$lib/types';
	let canvas: HTMLCanvasElement;
	const camera = new PerspectiveCamera(45, window.innerWidth / window.innerHeight, 0.1, 1000);
	const scene = new Scene();
	const loader = new GLTFLoader(new LoadingManager());
	const material = new MeshStandardMaterial({});
	const light = new AmbientLight(new Color('white'));
	const geometry = new BoxGeometry(1, 1, 1);
	const mesh = new Mesh(geometry, material);

	onMount(() => {
		loader.setWithCredentials(true);
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
		function render() {
			window.requestAnimationFrame(() => {
				controls.update();
				camera.updateProjectionMatrix();
				renderer.render(scene, camera);
				window.requestAnimationFrame(render);
			});
		}
		renderer.setSize(window.innerWidth, window.innerHeight);
		// scene.add(mesh);
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

<canvas bind:this={canvas} />
