import { useNavigate } from "react-router-dom"
import { Product } from "./Products"

type Props = {
    props:Product
    
}

const ProductCard = ({props}: Props) => {
    const navigate = useNavigate()
    const handleRedirect = () => {
        navigate(`/products/${props._id}`)
    }
  return (
      <div className="border border-black w-[23%] min-h-[50px] flex flex-col justify-center items-center p-4 gap-1" >
          <p>{props.name}</p>
          <div>
              <p>{props.brand}</p>
              <p>{props.content}</p>
              <p>{props.added_by}</p>
              <input type='submit' className='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded w--30' value="Explore" onClick={handleRedirect}/>
              
          </div>
    </div>
  )
}

export default ProductCard