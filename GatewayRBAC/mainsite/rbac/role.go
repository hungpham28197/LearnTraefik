package rbac

import "mainsite/pmodel"

func AnyRoles() RoleExp {
	return func() pmodel.Roles {
		mapRoles := make(pmodel.Roles)
		for _, role := range allRoles {
			mapRoles[role] = true
		}
		return mapRoles
	}
}

//Danh sách các role có thể truy xuất
func InRoles(roles ...int) RoleExp {
	return func() pmodel.Roles {
		mapRoles := make(pmodel.Roles)
		for _, role := range roles {
			mapRoles[role] = true
		}
		return mapRoles
	}
}

//Trả về map các role không nằm trong notRoles
func NotRoles(notRoles ...int) RoleExp {
	return func() pmodel.Roles {
		mapRoles := make(pmodel.Roles)
		for _, role := range allRoles {
			shouldInclude := true
			for _, notRole := range notRoles {
				if role == notRole { //Nếu role có trong danh sách NotRole thì không cho vào map
					shouldInclude = false
					break
				}
			}
			if shouldInclude { //Nếu role không có trong Not Role thì cho vào map
				mapRoles[role] = true
			}

		}
		return mapRoles
	}
}
