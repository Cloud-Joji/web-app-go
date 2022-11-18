import './App.css'

function App() {

  var baseUrl = window.location.href
  
  return (
    <div>
      <h1>Hello World with Vite & Cloud Run!</h1>
      <button onClick={ async () => {
        const response = await fetch(baseUrl + '/users')
        const data = await response.json()
        console.log(data)
        console.log(baseUrl)
      }}>Get Data</button>
    </div>
  )
}

export default App
