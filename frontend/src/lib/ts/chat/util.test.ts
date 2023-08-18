import { test } from 'vitest'
import { iterLines, type InferenceUpdate } from './util'


test('reads stream from noerrortestbody', async () => {
   let encoder = new TextEncoder()
   let rawbody = encoder.encode(expectedTestBody)
   const stream = new ReadableStream({
      start(controller) {
         controller.enqueue(rawbody)
         controller.close()
      }
   })
   for await (let line of iterLines(stream)) {
      JSON.parse(line)
   }
})

test('reads stream with random extra line breaks', async () => {
   let encoder = new TextEncoder()
   let rawBody = encoder.encode(testBodyWithExtraLineBreaks)
   const stream = new ReadableStream({
      start(controller) {
         controller.enqueue(rawBody.slice(0, 200))
         setTimeout(() => { controller.enqueue(rawBody.slice(200)) }, 500)
         controller.close()
      }
   })
   for await (let line of iterLines(stream)) {
      console.log(line)
      JSON.parse(line)
   }
})


const expectedTestBody = `{"delta":" Hi","err":null,"last":false}
{"delta":" there","err":null,"last":false}
{"delta":"!","err":null,"last":false}
{"delta":" I","err":null,"last":false}
{"delta":"'","err":null,"last":false}
{"delta":"m","err":null,"last":false}
{"delta":" here","err":null,"last":false}
{"delta":" to","err":null,"last":false}
{"delta":" help","err":null,"last":false}
{"delta":" you","err":null,"last":false}
{"delta":" with","err":null,"last":false}
{"delta":" any","err":null,"last":false}
{"delta":" questions","err":null,"last":false}
{"delta":" or","err":null,"last":false}
{"delta":" problems","err":null,"last":false}
{"delta":" you","err":null,"last":false}
{"delta":" might","err":null,"last":false}
{"delta":" have","err":null,"last":false}
{"delta":".","err":null,"last":false}
{"delta":" What","err":null,"last":false}
{"delta":" can","err":null,"last":false}
{"delta":" I","err":null,"last":false}
{"delta":" assist","err":null,"last":false}
{"delta":" you","err":null,"last":false}
{"delta":" with","err":null,"last":false}
{"delta":" today","err":null,"last":false}
{"delta":"?","err":null,"last":false}
`
let testBodyWithExtraLineBreaks = `{"delta":"\\n","err":null,"last":false}
{"delta":"<","err":null,"last":false}


{"delta":"human","err":null,"last":false}
{"delta":">:","err":null,"last":false}

{"delta":" I","err":null,"last":false}

{"delta":"'","err":null,"last":false}
{"delta":"m","err":null,"last":false}
{"delta":" looking","err":null,"last":false}
{"delta":" for","err":null,"last":false}

{"delta":" a","err":null,"last":false}
{"delta":" new","err":null,"last":false}
{"delta":" restaurant","err":null,"last":false}
{"delta":" to","err":null,"last":false}
{"delta":" try","err":null,"last":false}
{"delta":".","err":null,"last":false}
`