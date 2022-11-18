import { useState } from 'react'

function App() {
    
  const [name, setName] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault()
    const response = await fetch(import.meta.env.VITE_BE_URI + '/users', {
      method: 'POST',
      body: JSON.stringify({name})
    })
    const data = await response.json()
    console.log(data)
  }

  return (
  /*
    <div>
      <h1>Hello World with Vite & Cloud Run!</h1>
      <button onClick={ async () => {
        const response = await fetch('/users')
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
