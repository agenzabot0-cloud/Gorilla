import * as THREE from 'https://unpkg.com';

// 1. Core Setup
const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
const renderer = new THREE.WebGLRenderer({ antialias: true });

// 2. Configure Renderer and Inject into DOM
renderer.setSize(window.innerWidth, window.innerHeight);
document.body.style.margin = '0'; // Clean up page margins
document.body.appendChild(renderer.domElement);

// 3. Create the 3D Shape
const geometry = new THREE.BoxGeometry(1, 1, 1);
const material = new THREE.MeshStandardMaterial({ color: 0x3498db }); // Realistic blue material
const cube = new THREE.Mesh(geometry, material);
scene.add(cube);

// 4. Add Lighting (Necessary for MeshStandardMaterial)
const ambientLight = new THREE.AmbientLight(0xffffff, 0.5);
scene.add(ambientLight);

const directionalLight = new THREE.DirectionalLight(0xffffff, 1);
directionalLight.position.set(5, 5, 5);
scene.add(directionalLight);

camera.position.z = 3;

// 5. Handle Window Resizing Automatically
window.addEventListener('resize', () => {
    camera.aspect = window.innerWidth / window.innerHeight;
    camera.updateProjectionMatrix();
    renderer.setSize(window.innerWidth, window.innerHeight);
});

// 6. Animation Loop
function animate() {
    requestAnimationFrame(animate);
    
    // Rotate object
    cube.rotation.x += 0.01;
    cube.rotation.y += 0.01;
    
    renderer.render(scene, camera);
}

animate();
