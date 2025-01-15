import { Handle, NodeProps, Position } from '@xyflow/react'
import { FC } from 'react'

const NodeEnd: FC<NodeProps> = () => {
  return (
    <>
      <Handle type="target" position={Position.Left} />
      <div className="bg-[#474e93] px-[20px] py-[10px] rounded-md text-white" style={{ color: 'white' }}>
        End
      </div>
    </>
  )
}

export default NodeEnd
