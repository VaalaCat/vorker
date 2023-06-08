import ColorHash from 'color-hash'

const saturation = Array.from({ length: 45 }, (_, i) => i * 0.01)
export const CH = new ColorHash({ saturation: saturation })
