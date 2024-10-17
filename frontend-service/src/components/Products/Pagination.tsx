import React, { Dispatch, SetStateAction } from 'react'

type Props = {
    totalPages: number
    currentPage: number
    onPageChange:Dispatch<SetStateAction<number>>
}

const Pagination = ({ totalPages, currentPage, onPageChange }: Props) => {

  return (
      <div className='w-full bg-blue-300 flex justify-center gap-1'>
          {Array.from({ length: totalPages }, (_, i) => (
              <button className='border bg-red-200 p-2'
                  
              key={ i+ 1}
                  onClick={() => onPageChange(i + 1)}
              disabled={currentPage===i+1}
              >
                  {i+1}
            </button>
          ))}
    
         
      </div>
  )
}

export default Pagination