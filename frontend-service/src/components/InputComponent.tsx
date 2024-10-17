type Props = {
    name: string
    placeholder: string
    labelText?: string
  type?: string
  value:string
  onChange: (e:React.ChangeEvent<HTMLInputElement>)=> void;
}

const InputComponent = ({name,placeholder,labelText,type,value,onChange}:Props) => {
  return (
    <div className='w-1/2 flex flex-col gap-0.5 justify-between p-1 '  >
              <label className="capitalize" htmlFor={name}>{labelText||name}</label>
              <input className="p-1" type={type|| "text"} name={name} id={name} placeholder={placeholder} onChange={onChange} value={value} />
          </div>
  )
}

export default InputComponent