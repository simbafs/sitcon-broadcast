import { createRoot } from 'react-dom/client'
import { App } from './App'
import './main.css'

createRoot(document.getElementById('root')!).render(
	<div className="flex flex-col items-center justify-center h-screen w-screen">
		<App />
	</div>,
)
