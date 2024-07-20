import {defineConfig, loadEnv} from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {

  const env = loadEnv(mode, process.cwd(), '')
  const basePath = env.LOCAL_TO_DO_MANAGER_FRONTEND_BASE_PATH;
  console.log('FRONTEND_BASE_PATH: ', basePath)
  return {
    base: basePath,
    build: {
      outDir: '../dist'
    },
    plugins: [react()],
  };
})