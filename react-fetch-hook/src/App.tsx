import React from 'react';
import './App.css';

function App() {
  const [id, setId] = React.useState("")
  const status = useStatus(id);

  function onChange(e: React.ChangeEvent<HTMLInputElement>) {
    setId(e.target.value);
  }
  console.log('render', status);

  return (
    <div>
      <input type="text" onChange={(e) => onChange(e)} value={id} />
      <p>Status: {JSON.stringify(status)}</p>
    </div>
  );
}

function useStatus(id: string): string | null {
  const [state, setState] = React.useState<{[k: string]: string}>({});

  React.useEffect(() => {
    if (id === "") {
      return;
    }

    const timeout = setTimeout(() => {
      setState({[id]: id + ' is cool'})
    }, 1000);
    return () => clearTimeout(timeout);
  }, [id]);

  return (id === "")
    ? null
    : state[id] || 'loading';
}

export default App;
