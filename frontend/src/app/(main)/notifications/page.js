"use client";

import FetchFromBackend from "@/lib/fetch";
import fetchCredential from "@/lib/fetchCredential";
import Fetchnickname from "@/lib/fetchNickName";
import WsClient, { FollowRequest } from "@/lib/wsClient";
import { useState, useEffect, useRef } from "react";

export default function Notifications() {
  const [notifications, setNotifications] = useState([]);
  const [nickname, setNickname] = useState({})
  const ws = useRef(null);

  useEffect(() => {
    const load = async () => {
      try {
        const wsClient = await WsClient();
        ws.current = wsClient;

        ws.current.onmessage = (event) => {
          const data = JSON.parse(event.data)
          if (data.type === "follow_request") {
            alert(data.payload)
            setNotifications([...notifications, data.payload])
          }
        };
      } catch (error) {
        console.error(error);
      }
    };
    load();
  }, []);

  useEffect(() => {
    const GetNoti = async () => {
      try {
      const res = await FetchFromBackend("/notifications", {
        credentials: "include",
      })
      const data = await res.json()
      setNotifications([...notifications, data])
      } catch(err) {
        console.error(err)
      }
    }
    GetNoti()
  }, [])

  useEffect(() => {
    const findNickname = async () => {
      const ids = [...new Set(notifications.flatMap(followRequest => 
        followRequest.map(obj => obj.fromUserId)
      ))];
      console.log("these are the ids in notifications", ids);
      const nicknameMap = {};

      for (const userId of ids) {
        try {
          const res = await Fetchnickname(userId);
          const textData = await res.text();
          nicknameMap[userId] = textData;
          console.log(`nickname for id: ${userId} is ${textData}`)
        } catch (err) {
          console.error("error fetching nickanems for userids in notifications")
        }
      }
      console.log(nicknameMap)
      setNickname(nicknameMap)
    }
      findNickname()
  }, [notifications])

  const handleAccept = async (e) => {
    if (ws.current) {
      const fromUserId = e.target.getAttribute("from")
      const toUserId = e.target.getAttribute("to")
      const followsBack = true
      ws.current.send(
        JSON.stringify({
          type: "follow_request_status",
          payload: new FollowRequest(fromUserId, toUserId, followsBack),
        })
      )
      const notiBox = e.target.parentElement.parentElement;
      notiBox.remove()
    }
  };

  const handleDecline = (e) => {
    if (ws.current) {
      const fromUserId = e.target.getAttribute("from");
      const toUserId = e.target.getAttribute("to");
      const followsBack = false;
      ws.current.send(
        JSON.stringify({
          type: "follow_request_status",
          payload: new FollowRequest(fromUserId, toUserId, followsBack),
        })
      )
      const notiBox = e.target.parentElement.parentElement;
      notiBox.remove()
    }
  };

  return (
    <div>
      <h1>Notifications</h1>
      <section>
        {notifications.map((followRequest) => {
          return followRequest.map((obj, i) => {
            return (
              <article className="bg-secondary p-4 rounded-lg m-4 flex justify-between w-[30vw]" id={obj.toUserId} key={i}>
                <main>
                  <h1>FollowRequest from {nickname[obj.fromUserId]}</h1>
                </main>
                <aside>
                  <button className="bg-green-700 ml-4 rounded-lg p-2 text-white"  onClick={handleAccept} from={obj.fromUserId} to={obj.toUserId} key="Accept">accept</button>
                  <button className="bg-red-700 ml-2 rounded-lg p-2 text-white" onClick={handleDecline} from={obj.fromUserId} to={obj.toUserId} key="Decline">decline</button>
                </aside>
              </article>
            )
          })
        })}
      </section>
    </div>
  );
}
