
'use client'

import { useEffect, useState } from 'react'



export default function Home() {
  const [start, setStart] = useState('')
  const [end, setEnd] = useState('')
  const [data,setData] = useState([])
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    fetchData() 
  }, [start, end])

  async function fetchData() {
    setLoading(true)

    const params = new URLSearchParams()
    if (start) params.append('start', start)
    if (end) params.append('end', end)


    const res = await fetch(
      `http://localhost:8080/api/data?${params.toString()}`
    )
    const data = await res.json()
    setData(Array.isArray(data) ? data : [])
    setLoading(false)
  }
  return (
    <div className="flex min-h-screen items-center justify-center bg-zinc-50 font-sans dark:bg-black">
      <h1>Vehicles Trips</h1>
        <div>
          <input
            type="date"
            className="border rounded-md px-3 py-2"
            value={start}
            onChange={e => setStart(e.target.value)}
          />

          <input
            type="date"
            className="border rounded-md px-3 py-2"
            value={end}
            onChange={e => setEnd(e.target.value)}
          />
        </div>

         <div className="overflow-x-auto">
          <table className="w-full border border-gray-200">
            <thead className="bg-gray-50">
              <tr>
                <th className="text-left px-4 py-2 border">Vehicle ID</th>
                <th className="text-left px-4 py-2 border">Driver Name</th>
                <th className="text-left px-4 py-2 border">Start</th>
                <th className="text-left px-4 py-2 border">End</th>
              </tr>
            </thead>
            <tbody>
              {loading ? (
                <tr>
                  <td colSpan={4} className="text-center py-6">
                    Loading...
                  </td>
                </tr>
              ) : data.length === 0 ? (
                <tr>
                  <td colSpan={4} className="text-center py-6 text-gray-500">
                    No data found
                  </td>
                </tr>
              ) : (
                data.map(a => (
                  <tr
                    key={a.id}
                    className="hover:bg-gray-50 transition"
                  >
                    <td className="px-4 py-2 border">{a.vehicleID}</td>
                    <td className="px-4 py-2 border">{a.driverName}</td>
                    <td className="px-4 py-2 border">{a.startDate}</td>
                    <td className="px-4 py-2 border">{a.endDate}</td>
                    <td className="px-4 py-2 border">
                    </td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>
    </div>
  );
}