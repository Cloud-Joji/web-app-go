import { useState } from 'react'

function App() {
    
  const [name, setName] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault()
    const response = await fetch('/certs', {
      method: 'POST',
      body: JSON.stringify({name}),
      headers: {
        "Content-Type" : "application/json"
      }
    })
    const data = await response.json()
    console.log(data)
  }

  return (
  /*
    <div>
      <h1>Hello World with Vite & Cloud Run!</h1>
      <button onClick={ async () => {
        const response = await fetch('/certs')
        const data = await response.json()
        console.log(data)
      }}>Get Data</button>
    </div>
  */

    <div>

    <form onSubmit={handleSubmit}>
      <input 
        type="name" 
        placeholder="Certification name" 
        onChange={e => setName(e.target.value)} 
      />

      <button>Save</button>
    </form>

    </div>
  )
}

export default App
