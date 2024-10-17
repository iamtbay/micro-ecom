
import { BrowserRouter as Router,Route,Routes } from 'react-router-dom'
import Navbar from './components/Navbar'
import Mainpage from './pages/Mainpage'
import AuthPage from './pages/AuthPage'
import ProductsPage from './pages/ProductsPage'

function App() {
  return (
    <>
    <Navbar />
    <Router>
      <Routes>
        <Route path='/' element={<Mainpage/>} />
        <Route path='/auth' element={<AuthPage/>} />
        <Route path='/products' element={<ProductsPage/>} />
      </Routes>
    </Router>
    </>
  
  )
}

export default App
