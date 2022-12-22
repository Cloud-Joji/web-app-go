import { useEffect, useState } from 'react'

function App() {
    
  const [certs, setCerts] = useState([])
  const [name, setName] = useState('');
  const [platform, setPlatform] = useState('Choose')

  async function loadCerts(){
    const response = await fetch('/api/certs')  
    const data = await response.json()
    setCerts(data.certs)
  }

  useEffect(() => {
    loadCerts()
  }, [])

  const handleSubmit = async (e) => {
    e.preventDefault()
    console.log(name)
    console.log(platform)
    console.log(JSON.stringify({name, platform}))
    const response = await fetch('/api/certs', {
      method: 'POST',
      body: JSON.stringify({name, platform}),
      headers: {
        "Content-Type" : "application/json"
      }
    })
    const data = await response.json()
    loadCerts()
  }

  return (
    <div>

    <form onSubmit={handleSubmit}>
      <label>
        Certification Name 
        <input 
          type="name" 
          placeholder="Certification name" 
          onChange={e => setName(e.target.value)} 
        />
      </label>

      <label>
        Certification Platform 
      <select value={platform} onChange={e => setPlatform(e.target.value)}>
        <option value="Codigo Facilito">Codigo Facilito</option>
        <option value="Coursera">Coursera</option>
        <option value="ED Team">ED Team</option>
        <option value="LinkedIn Learning">LinkedIn Learning</option>
        <option value="Platzi">Platzi</option>
        <option value="Udemy">Udemy</option>
        <option value="Choose">Choose</option>
      </select>
      </label>

      <button>Save</button>
    </form>

    <ul>
      {certs.map(cert => (
        <li key={cert._id}>{cert.name} @ {cert.platform}</li>
      ))}
    </ul>

    </div>
  )
}

export default App
