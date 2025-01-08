package db

import (
	"fmt"
	"log"
	"social-network/pkg/helpers"
	"strings"
)

func GetGroupFromDb(groupname string) (helpers.Group, error) {
	group := helpers.Group{}

	rows, err := DB.Query("SELECT * FROM groups WHERE title = ?", groupname)
	if err != nil {
		log.Println("Query error:", err)
		return group, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&group.CreatorName, &group.Title, &group.Description)
		if err != nil {
			log.Println("Scan error:", err)
			return group, err
		}
	}

	return group, nil
}

func GetGroupsFromDb() ([]helpers.Group, error) {
	groups := []helpers.Group{}

	rows, err := DB.Query("SELECT * FROM groups")
	if err != nil {
		log.Println("Query error:", err)
		return groups, err
	}
	defer rows.Close()
	for rows.Next() {
		group := helpers.Group{}
		err := rows.Scan(&group.CreatorName, &group.Title, &group.Description)
		if err != nil {
			log.Println("Scan error:", err)
			return groups, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

func CreateGroupDB(data []interface{}) error {

	stmt, err := DB.Prepare("INSERT INTO groups (creator_name, title, description) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("Prepare statement error:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data...)
	if err != nil {
		log.Println("Exec statement error:", err)
		return err
	}

	return nil
}

type MembershipExistsError struct {
	Status string
}

func (e *MembershipExistsError) Error() string {
	return fmt.Sprintf("Membership already exists with status: %s", e.Status)
}

func InviteMemberDB(groupname string, username string) (string, bool) {

	var existingStatus string
	err := DB.QueryRow("SELECT status FROM memberships WHERE title = ? AND nickname = ?", groupname, username).Scan(&existingStatus)
	log.Println("TEST", existingStatus, err)

	// already there is data for the user & the group
	if existingStatus != "" {
		log.Println("Existing status found:", existingStatus)

		switch existingStatus {
		case "requested":
			log.Println("Case: requested")
			return fmt.Sprintf(`User "%s" has already requested to join the group.`, username), true
		case "invited":
			log.Println("Case: invited")
			return fmt.Sprintf(`User "%s" has already been invited to the group.`, username), true
		case "approved":
			log.Println("Case: approved")
			return fmt.Sprintf(`User "%s" is already a member of the group.`, username), true
		default:
			log.Println("Unexpected status encountered")
			return "Unexpected status.", true
		}
	}

	// make new data
	stmt, err := DB.Prepare("INSERT INTO memberships (title, nickname, status) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("Prepare statement error:", err)
		return fmt.Sprintln("Prepare statement error:", err), true
	}
	defer stmt.Close()
	_, err = stmt.Exec(groupname, username, "invited")

	// the user does not exist
	if err != nil {
		if err.Error() == "FOREIGN KEY constraint failed" {
			return fmt.Sprintf(`Invitation unsent: user "%s" does not exist`, username), true
		}
		return fmt.Sprintln("Exac statement error:", err), true
	}

	return "", false
}

func UpdateMemberStatus(id int, groupname string, username string, status string) {
	log.Println("updating member stat... ", id, groupname, username, status)

	var err error

	if status == "requested" && id == 0 {
		log.Println("requested", id, groupname, username, status)
		query := `INSERT INTO memberships (title, nickname, status) VALUES (?, ?, ?)`
		_, err = DB.Exec(query, groupname, username, "requested")
	}

	if status == "approve" {
		log.Println("approved", id, groupname, username, status)
		query := `UPDATE memberships SET status = ? WHERE id = ?`
		_, err = DB.Exec(query, "approved", id)
	}

	if status == "reject" {
		log.Println("rejected", id, groupname, username, status)

		query := `DELETE FROM memberships WHERE id = ?`
		_, err = DB.Exec(query, id)
	}

	if err != nil {
		log.Printf("failed to update member status %v: ", status)
		log.Println(err)
	}
}

func GetMembershipsFromDb(nickname string) ([]helpers.Membership, error) {
	rows, err := DB.Query("SELECT * FROM memberships WHERE nickname = ?", nickname)
	if err != nil {
		log.Println("Query error:", err)
		return nil, err
	}

	defer rows.Close()
	memberships := []helpers.Membership{}
	for rows.Next() {
		membership := helpers.Membership{}
		err := rows.Scan(&membership.Id, &membership.Title, &membership.Username, &membership.Status)
		if err != nil {
			log.Println("Scan error:", err)
			return nil, err
		}
		memberships = append(memberships, membership)
	}

	log.Println(memberships)

	return memberships, nil
}

func AddGroupPostToDb(data []interface{}) error {
	log.Println("grouppost DB", data)
	stmt, err := DB.Prepare("INSERT INTO group_posts (group_title, user_id, nickname, subject, content, image) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Prepare statement error:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data...)
	if err != nil {
		log.Println("Exec statement error:", err)
		return err
	}
	return nil
}

func GetGroupPostsFromDb(groupname string) ([]helpers.GroupPost, error) {
	posts := []helpers.GroupPost{}

	rows, err := DB.Query("SELECT * FROM group_posts WHERE group_title = ?", groupname)
	if err != nil {
		log.Println("Query error:", err)
		return posts, err
	}
	defer rows.Close()
	for rows.Next() {
		post := helpers.GroupPost{}
		err := rows.Scan(&post.Id, &post.GroupTitle, &post.UserID, &post.Nickname, &post.Subject, &post.Content, &post.Image, &post.CreationDate)
		if err != nil {
			log.Println("Scan error:", err)
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetGroupPostFromDb(postID int) (helpers.GroupPost, error) {
	post := helpers.GroupPost{}

	rows, err := DB.Query("SELECT * FROM group_posts WHERE id = ?", postID)
	if err != nil {
		log.Println("Query error:", err)
		return post, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&post.Id, &post.GroupTitle, &post.UserID, &post.Nickname, &post.Subject, &post.Content, &post.Image, &post.CreationDate)
		if err != nil {
			log.Println("Scan error:", err)
			return post, err
		}
	}

	return post, nil
}

func GetGroupPostCommentsFromDb(postID int) ([]helpers.GroupComment, error) {
	comments := []helpers.GroupComment{}

	rows, err := DB.Query("SELECT * FROM group_comments WHERE post_id = ?", postID)
	if err != nil {
		log.Println("Query error:", err)
		return comments, err
	}
	defer rows.Close()
	for rows.Next() {
		comment := helpers.GroupComment{}
		err := rows.Scan(&comment.CommentId, &comment.PostId, &comment.UserID, &comment.Nickname, &comment.Content, &comment.Image, &comment.CreationDate)
		if err != nil {
			log.Println("Scan error:", err)
			return comments, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func AddGroupPostCommentToDb(data []interface{}) error {
	log.Println("comment DB", data)
	stmt, err := DB.Prepare("INSERT INTO group_comments (post_id, user_id, nickname, content, image) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Prepare statement error:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data...)
	if err != nil {
		log.Println("Exec statement error:", err)
		return err
	}
	return nil
}

func GetGroupEventsFromDb(groupname string) []helpers.GroupEvent {
	events := []helpers.GroupEvent{}

	rows, err := DB.Query("SELECT * FROM group_events WHERE group_title = ?", groupname)
	if err != nil {
		log.Println("Query error:", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		event := helpers.GroupEvent{}
		err := rows.Scan(&event.Id, &event.GroupTitle, &event.UserID, &event.Nickname, &event.Title, &event.Description, &event.EventDate)
		if err != nil {
			log.Println("Scan error:", err)
			return nil
		}
		events = append(events, event)
	}

	return events
}

func AddGroupEventToDb(data []interface{}) error {
	log.Println("groupevent DB", data)
	stmt, err := DB.Prepare("INSERT INTO group_events (group_title, user_id, nickname, title, description, event_date) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Prepare statement error:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data...)
	if err != nil {
		log.Println("Exec statement error:", err)
		return err
	}
	return nil
}

func GetEventStatusFromDb(nickname string, eventID int) helpers.EventStatus {
	var eventStatus helpers.EventStatus

	rows, err := DB.Query("SELECT * FROM user_event_status WHERE nickname = ? AND event_id = ?", nickname, eventID)
	if err != nil {
		log.Println("Query error:", err)
		return eventStatus
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&eventStatus.Id, &eventStatus.Nickname, &eventStatus.EventId, &eventStatus.Going)
		if err != nil {
			log.Println("Scan error:", err)
			return eventStatus
		}
	}
	log.Println(eventStatus)

	return eventStatus
}

func UpdateEventStatus(nickname string, EventId int, status bool) {
	query := `
	INSERT OR REPLACE INTO user_event_status (nickname, event_id, going)
	VALUES (?, ?, ?)
	`
	_, err := DB.Exec(query, nickname, EventId, status)
	if err != nil {
		log.Print("failed to insert or update event status: %w", err)
	}
}

func GetYourRequestsFromDb(groups []helpers.Membership) []helpers.Membership {
	// get titles
	titles := []string{}
	for _, group := range groups {
		titles = append(titles, group.Title)
	}

	placeholders := make([]string, len(titles))
	args := []interface{}{"requested"}
	for i, title := range titles {
		placeholders[i] = "?"
		args = append(args, title)
	}
	query := fmt.Sprintf("SELECT * FROM memberships WHERE status = ? AND title IN (%s)", strings.Join(placeholders, ", "))

	rows, err := DB.Query(query, args...)
	if err != nil {
		log.Println("Query error:", err)
		return nil
	}
	defer rows.Close()

	memberships := []helpers.Membership{}
	for rows.Next() {
		membership := helpers.Membership{}
		err := rows.Scan(&membership.Id, &membership.Title, &membership.Username, &membership.Status)
		if err != nil {
			log.Println("Scan error:", err)
			return nil
		}
		memberships = append(memberships, membership)
	}

	log.Println(memberships)
	return memberships
}
