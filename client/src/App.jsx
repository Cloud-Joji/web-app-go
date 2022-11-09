import './App.css'

function App() {

  return (
    <div>
      <h1>Hello World with Vite</h1>
      <button onClick={async () => {
        const response = await fetch('http://localhost:4000/users')
        const data = await response.json()
        console.log(data)
      }}>Get Data</button>
    </div>
  )
}

export default App
