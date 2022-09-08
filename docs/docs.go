// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "https://t.me/vmkevv",
            "email": "vargaskevv@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/deps": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "List all deps",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Dpto"
                            }
                        }
                    }
                }
            }
        },
        "/provs/dep/{depid}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "List all deps",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Departamento ID",
                        "name": "depid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Provincia"
                            }
                        }
                    }
                }
            }
        },
        "/signin": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "login and generate token",
                "parameters": [
                    {
                        "description": "teacher signup data",
                        "name": "teacher",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.SignInHandler.req"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.SignInHandler.res"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "registers a new teacher",
                "parameters": [
                    {
                        "description": "teacher signup data",
                        "name": "teacher",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.SignUpHandler.req"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.SignUpHandler.res"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ent.Activity": {
            "type": "object",
            "properties": {
                "date": {
                    "description": "Date holds the value of the \"date\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ActivityQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.ActivityEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "ent.ActivityEdges": {
            "type": "object",
            "properties": {
                "area": {
                    "description": "Area holds the value of the area edge.",
                    "$ref": "#/definitions/ent.Area"
                },
                "classPeriod": {
                    "description": "ClassPeriod holds the value of the classPeriod edge.",
                    "$ref": "#/definitions/ent.ClassPeriod"
                },
                "scoreSyncs": {
                    "description": "ScoreSyncs holds the value of the scoreSyncs edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.ScoreSync"
                    }
                },
                "scores": {
                    "description": "Scores holds the value of the scores edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Score"
                    }
                }
            }
        },
        "ent.ActivitySync": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ActivitySyncQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.ActivitySyncEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "last_sync_id": {
                    "description": "LastSyncID holds the value of the \"last_sync_id\" field.",
                    "type": "string"
                }
            }
        },
        "ent.ActivitySyncEdges": {
            "type": "object",
            "properties": {
                "classPeriod": {
                    "description": "ClassPeriod holds the value of the classPeriod edge.",
                    "$ref": "#/definitions/ent.ClassPeriod"
                }
            }
        },
        "ent.Area": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the AreaQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.AreaEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                },
                "points": {
                    "description": "Points holds the value of the \"points\" field.",
                    "type": "integer"
                }
            }
        },
        "ent.AreaEdges": {
            "type": "object",
            "properties": {
                "activities": {
                    "description": "Activities holds the value of the activities edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Activity"
                    }
                },
                "year": {
                    "description": "Year holds the value of the year edge.",
                    "$ref": "#/definitions/ent.Year"
                }
            }
        },
        "ent.Attendance": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the AttendanceQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.AttendanceEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "value": {
                    "description": "Value holds the value of the \"value\" field.",
                    "type": "string"
                }
            }
        },
        "ent.AttendanceEdges": {
            "type": "object",
            "properties": {
                "classPeriod": {
                    "description": "ClassPeriod holds the value of the classPeriod edge.",
                    "$ref": "#/definitions/ent.ClassPeriod"
                },
                "student": {
                    "description": "Student holds the value of the student edge.",
                    "$ref": "#/definitions/ent.Student"
                }
            }
        },
        "ent.AttendanceSync": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the AttendanceSyncQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.AttendanceSyncEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "last_sync_id": {
                    "description": "LastSyncID holds the value of the \"last_sync_id\" field.",
                    "type": "string"
                }
            }
        },
        "ent.AttendanceSyncEdges": {
            "type": "object",
            "properties": {
                "classPeriod": {
                    "description": "ClassPeriod holds the value of the classPeriod edge.",
                    "$ref": "#/definitions/ent.ClassPeriod"
                }
            }
        },
        "ent.Class": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ClassQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.ClassEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "parallel": {
                    "description": "Parallel holds the value of the \"parallel\" field.",
                    "type": "string"
                }
            }
        },
        "ent.ClassEdges": {
            "type": "object",
            "properties": {
                "classPeriodSyncs": {
                    "description": "ClassPeriodSyncs holds the value of the classPeriodSyncs edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.ClassPeriodSync"
                    }
                },
                "classPeriods": {
                    "description": "ClassPeriods holds the value of the classPeriods edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.ClassPeriod"
                    }
                },
                "grade": {
                    "description": "Grade holds the value of the grade edge.",
                    "$ref": "#/definitions/ent.Grade"
                },
                "school": {
                    "description": "School holds the value of the school edge.",
                    "$ref": "#/definitions/ent.School"
                },
                "studentSyncs": {
                    "description": "StudentSyncs holds the value of the studentSyncs edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.StudentSync"
                    }
                },
                "students": {
                    "description": "Students holds the value of the students edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Student"
                    }
                },
                "subject": {
                    "description": "Subject holds the value of the subject edge.",
                    "$ref": "#/definitions/ent.Subject"
                },
                "teacher": {
                    "description": "Teacher holds the value of the teacher edge.",
                    "$ref": "#/definitions/ent.Teacher"
                }
            }
        },
        "ent.ClassPeriod": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ClassPeriodQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.ClassPeriodEdges"
                },
                "end": {
                    "description": "End holds the value of the \"end\" field.",
                    "type": "string"
                },
                "finished": {
                    "description": "Finished holds the value of the \"finished\" field.",
                    "type": "boolean"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "start": {
                    "description": "Start holds the value of the \"start\" field.",
                    "type": "string"
                }
            }
        },
        "ent.ClassPeriodEdges": {
            "type": "object",
            "properties": {
                "activities": {
                    "description": "Activities holds the value of the activities edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Activity"
                    }
                },
                "activitySyncs": {
                    "description": "ActivitySyncs holds the value of the activitySyncs edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.ActivitySync"
                    }
                },
                "attendanceSyncs": {
                    "description": "AttendanceSyncs holds the value of the attendanceSyncs edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.AttendanceSync"
                    }
                },
                "attendances": {
                    "description": "Attendances holds the value of the attendances edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Attendance"
                    }
                },
                "class": {
                    "description": "Class holds the value of the class edge.",
                    "$ref": "#/definitions/ent.Class"
                },
                "period": {
                    "description": "Period holds the value of the period edge.",
                    "$ref": "#/definitions/ent.Period"
                }
            }
        },
        "ent.ClassPeriodSync": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ClassPeriodSyncQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.ClassPeriodSyncEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "last_sync_id": {
                    "description": "LastSyncID holds the value of the \"last_sync_id\" field.",
                    "type": "string"
                }
            }
        },
        "ent.ClassPeriodSyncEdges": {
            "type": "object",
            "properties": {
                "class": {
                    "description": "Class holds the value of the class edge.",
                    "$ref": "#/definitions/ent.Class"
                }
            }
        },
        "ent.Dpto": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the DptoQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.DptoEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "ent.DptoEdges": {
            "type": "object",
            "properties": {
                "provincias": {
                    "description": "Provincias holds the value of the provincias edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Provincia"
                    }
                }
            }
        },
        "ent.Grade": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the GradeQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.GradeEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "ent.GradeEdges": {
            "type": "object",
            "properties": {
                "classes": {
                    "description": "Classes holds the value of the classes edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Class"
                    }
                }
            }
        },
        "ent.Municipio": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the MunicipioQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.MunicipioEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "ent.MunicipioEdges": {
            "type": "object",
            "properties": {
                "provincia": {
                    "description": "Provincia holds the value of the provincia edge.",
                    "$ref": "#/definitions/ent.Provincia"
                },
                "schools": {
                    "description": "Schools holds the value of the schools edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.School"
                    }
                }
            }
        },
        "ent.Period": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the PeriodQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.PeriodEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "ent.PeriodEdges": {
            "type": "object",
            "properties": {
                "classPeriods": {
                    "description": "ClassPeriods holds the value of the classPeriods edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.ClassPeriod"
                    }
                },
                "year": {
                    "description": "Year holds the value of the year edge.",
                    "$ref": "#/definitions/ent.Year"
                }
            }
        },
        "ent.Provincia": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ProvinciaQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.ProvinciaEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "ent.ProvinciaEdges": {
            "type": "object",
            "properties": {
                "departamento": {
                    "description": "Departamento holds the value of the departamento edge.",
                    "$ref": "#/definitions/ent.Dpto"
                },
                "municipios": {
                    "description": "Municipios holds the value of the municipios edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Municipio"
                    }
                }
            }
        },
        "ent.School": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the SchoolQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.SchoolEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "lat": {
                    "description": "Lat holds the value of the \"lat\" field.",
                    "type": "string"
                },
                "lon": {
                    "description": "Lon holds the value of the \"lon\" field.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "ent.SchoolEdges": {
            "type": "object",
            "properties": {
                "classes": {
                    "description": "Classes holds the value of the classes edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Class"
                    }
                },
                "municipio": {
                    "description": "Municipio holds the value of the municipio edge.",
                    "$ref": "#/definitions/ent.Municipio"
                }
            }
        },
        "ent.Score": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ScoreQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.ScoreEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "points": {
                    "description": "Points holds the value of the \"points\" field.",
                    "type": "integer"
                }
            }
        },
        "ent.ScoreEdges": {
            "type": "object",
            "properties": {
                "activity": {
                    "description": "Activity holds the value of the activity edge.",
                    "$ref": "#/definitions/ent.Activity"
                },
                "student": {
                    "description": "Student holds the value of the student edge.",
                    "$ref": "#/definitions/ent.Student"
                }
            }
        },
        "ent.ScoreSync": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ScoreSyncQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.ScoreSyncEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "last_sync_id": {
                    "description": "LastSyncID holds the value of the \"last_sync_id\" field.",
                    "type": "string"
                }
            }
        },
        "ent.ScoreSyncEdges": {
            "type": "object",
            "properties": {
                "activity": {
                    "description": "Activity holds the value of the activity edge.",
                    "$ref": "#/definitions/ent.Activity"
                }
            }
        },
        "ent.Student": {
            "type": "object",
            "properties": {
                "ci": {
                    "description": "Ci holds the value of the \"ci\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the StudentQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.StudentEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "last_name": {
                    "description": "LastName holds the value of the \"last_name\" field.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "ent.StudentEdges": {
            "type": "object",
            "properties": {
                "attendances": {
                    "description": "Attendances holds the value of the attendances edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Attendance"
                    }
                },
                "class": {
                    "description": "Class holds the value of the class edge.",
                    "$ref": "#/definitions/ent.Class"
                },
                "scores": {
                    "description": "Scores holds the value of the scores edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Score"
                    }
                }
            }
        },
        "ent.StudentSync": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the StudentSyncQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.StudentSyncEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "last_sync_id": {
                    "description": "LastSyncID holds the value of the \"last_sync_id\" field.",
                    "type": "string"
                }
            }
        },
        "ent.StudentSyncEdges": {
            "type": "object",
            "properties": {
                "class": {
                    "description": "Class holds the value of the class edge.",
                    "$ref": "#/definitions/ent.Class"
                }
            }
        },
        "ent.Subject": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the SubjectQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.SubjectEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "ent.SubjectEdges": {
            "type": "object",
            "properties": {
                "classes": {
                    "description": "Classes holds the value of the classes edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Class"
                    }
                }
            }
        },
        "ent.Teacher": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the TeacherQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.TeacherEdges"
                },
                "email": {
                    "description": "Email holds the value of the \"email\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "last_name": {
                    "description": "LastName holds the value of the \"last_name\" field.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                },
                "password": {
                    "description": "Password holds the value of the \"password\" field.",
                    "type": "string"
                }
            }
        },
        "ent.TeacherEdges": {
            "type": "object",
            "properties": {
                "classes": {
                    "description": "Classes holds the value of the classes edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Class"
                    }
                }
            }
        },
        "ent.Year": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the YearQuery when eager-loading is set.",
                    "$ref": "#/definitions/ent.YearEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "value": {
                    "description": "Value holds the value of the \"value\" field.",
                    "type": "integer"
                }
            }
        },
        "ent.YearEdges": {
            "type": "object",
            "properties": {
                "areas": {
                    "description": "Areas holds the value of the areas edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Area"
                    }
                },
                "classes": {
                    "description": "Classes holds the value of the classes edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Class"
                    }
                },
                "periods": {
                    "description": "Periods holds the value of the periods edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Period"
                    }
                }
            }
        },
        "handlers.SignInHandler.req": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "handlers.SignInHandler.res": {
            "type": "object",
            "properties": {
                "jwt": {
                    "type": "string"
                },
                "teacher": {
                    "$ref": "#/definitions/ent.Teacher"
                }
            }
        },
        "handlers.SignUpHandler.req": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "handlers.SignUpHandler.res": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:4000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Rigel API",
	Description:      "Backend json services for Rigel WebApp",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
