let ws = null;

export default async function WsClient() {
  if (
    !ws ||
    ws.readyState === WebSocket.CLOSING ||
    ws.readyState === WebSocket.CLOSED
  ) {
    const wsUrl = process.env.NEXT_PUBLIC_WS_URL;
    ws = new WebSocket(wsUrl);

    ws.onopen = async () => {
      console.log("ws connected");
    };

    ws.onclose = () => {
      console.log("ws disconnected");
    };

    ws.onerror = (err) => {
      console.error("ws error", err);
    };
  }
  return ws;
}

class Event {
  constructor(type, data) {
    this.type = type;
    this.data = data;
  }
}

// Add diferent event types down here
// these will be sent to the backend
export class FollowRequest {
  constructor(from, to, followsBack) {
    this.fromUserId = from; // user id
    this.toUserId = to; // user id
    this.followsBack = followsBack;
  }
}

export class ChatMessage {
  constructor(chatType, from, groupId, content) {
    this.chatType = chatType
    this.groupId = groupId
    this.fromUserId = from
    this.content = content
  }
}
